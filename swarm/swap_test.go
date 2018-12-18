// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package swarm

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/simulations/adapters"
	"github.com/ethereum/go-ethereum/swarm/api"
	"github.com/ethereum/go-ethereum/swarm/log"
	"github.com/ethereum/go-ethereum/swarm/network/simulation"
	"github.com/ethereum/go-ethereum/swarm/storage"
)

var (
	printStats     = flag.Bool("printstats", false, "print swap stats")
	bucketKeySwarm = simulation.BucketKey("swarm")
)

//In TestSwapNetworkSymmetricFileUpload we set up a network with arbitrary number of nodes
//(16), and each of the nodes uploads a file of same size
//Afterwards we check that every node's balance WITH ANOTHER PEER
//has the same value but opposite sign
func TestSwapNetworkSymmetricFileUpload(t *testing.T) {
	//default hardcoded network size
	nodeCount := 16
	//every node has a map to all nodes it had interactions
	//each entry in the map is a map of the other node with all the balances
	balancesMap := make(map[enode.ID]map[enode.ID]int64)

	//setup the simulation
	//use a complete node setup via `NewSwam`
	sim := simulation.New(map[string]simulation.ServiceFunc{
		"swarm": func(ctx *adapters.ServiceContext, bucket *sync.Map) (s node.Service, cleanup func(), err error) {
			config := api.NewConfig()
			config.Port = ""

			dir, err := ioutil.TempDir("", "swap-network-test-node")
			if err != nil {
				return nil, nil, err
			}
			cleanup = func() {
				err := os.RemoveAll(dir)
				if err != nil {
					log.Error("cleaning up swarm temp dir", "err", err)
				}
			}

			config.Path = dir

			privkey, err := crypto.GenerateKey()
			if err != nil {
				return nil, cleanup, err
			}

			config.Init(privkey)

			//set Swap to be enabled for this test
			config.SwapEnabled = true

			swarm, err := NewSwarm(config, nil)
			if err != nil {
				return nil, cleanup, err
			}

			bucket.Store(bucketKeySwarm, swarm)
			log.Info("new swarm", "bzzKey", config.BzzKey, "baseAddr", fmt.Sprintf("%x", swarm.bzz.BaseAddr()))
			return swarm, cleanup, nil
		},
	})

	ctx := context.Background()
	files := make([]file, 0)

	//upload a snapshot
	err := sim.UploadSnapshot(fmt.Sprintf("network/stream/testing/snapshot_%d.json", nodeCount))
	if err != nil {
		t.Fatal(err)
	}
	//run the simulation
	result := sim.Run(ctx, func(ctx context.Context, sim *simulation.Simulation) error {
		//wait for kademlia to be healthy
		if _, err := sim.WaitTillHealthy(ctx, 2); err != nil {
			return err
		}

		nodeIDs := sim.UpNodeIDs()
		rand.Shuffle(len(nodeIDs), func(i, j int) {
			nodeIDs[i], nodeIDs[j] = nodeIDs[j], nodeIDs[i]
		})
		//upload a file for every node
		for _, id := range nodeIDs {
			item, ok := sim.NodeItem(id, bucketKeySwarm)
			if !ok {
				return errors.New("No swarm")
			}
			swarm := item.(*Swarm)
			key, data, err := uploadFile(swarm)
			if err != nil {
				return err
			}
			log.Trace("file uploaded", "node", id, "key", key.String())
			files = append(files, file{
				addr:   key,
				data:   data,
				nodeID: id,
			})
		}

		// File retrieval check is repeated until all uploaded files are retrieved from all nodes
		// or until the timeout is reached.
		//we use a special retrieve function for swap which is optimized for parallel requests
		//but does not leave many cascaded requests floating around
		for retrieveForSwap(sim, files) != 0 {
		}

		//iterate all nodes
		for _, node := range sim.NodeIDs() {
			item, ok := sim.NodeItem(node, bucketKeySwarm)
			if !ok {
				return errors.New("No swarm")
			}
			swarm := item.(*Swarm)

			//submap for each node is a map of all nodes with the balance for that node
			subBalances := make(map[enode.ID]int64)

			//iterate all nodes again...
			//get all balances with other peers for every node
			for _, n := range sim.NodeIDs() {
				if node == n {
					continue
				}

				//get the peer's balance with this node
				balance, err := swarm.swap.GetPeerBalance(n)
				if err == nil {
					subBalances[n] = balance
					log.Debug(fmt.Sprintf("Balance of node %s to node %s: %d", node.TerminalString(), n.TerminalString(), balance))
				} else {
					log.Debug(fmt.Sprintf("Node %s has no balance with node %s", node.TerminalString(), n.TerminalString()))
				}
			}
			//update the map for this node
			balancesMap[node] = subBalances
		}

		return nil
	})

	sim.Close()
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	//print all the balances if requested
	if *printStats {
		for k, v := range balancesMap {
			fmt.Println(fmt.Sprintf("node %s balances:", k.TerminalString()))
			for kk, vv := range v {
				fmt.Println(fmt.Sprintf(".........with node %s: balance %d", kk.TerminalString(), vv))
			}
		}
	}

	//now iterate the whole map
	//and check that every node k has the same
	//balance with a peer as that peer with the node,
	//but in inverted signs

	//iterate the map
	errorFound := false
	for k, mapForK := range balancesMap {
		//iterate the submap
		for n, balanceKwithN := range mapForK {
			//iterate the main map again
			mapForSubK := balancesMap[n]
			log.Trace(fmt.Sprintf("balance of %s with %s: %d", k.TerminalString(), n.TerminalString(), balanceKwithN))
			log.Trace(fmt.Sprintf("balance of %s with %s: %d", n.TerminalString(), k.TerminalString(), mapForSubK[k]))
			//...check that they have the same balance in Abs terms and that it is not 0
			if balanceKwithN+mapForSubK[k] != 0 && balanceKwithN != 0 {
				log.Error("Expected balances to be a+b = 0 AND balance(a) != 0, but they are not", "balance_k_with", balanceKwithN, "balance_n_with_k", mapForSubK[k])
				errorFound = true
			}
		}
	}
	if errorFound {
		t.Fatal("Expected balances to be symmetrical, but they were not")
	}
	log.Debug("test terminated")
}

//TestSwapNetworkAsymmetricFileUpload is a swap test too,
//but this time the number and size of files are random
func TestSwapNetworkAsymmetricFileUpload(t *testing.T) {
	nodeCount := 16
	balancesMap := make(map[enode.ID]map[enode.ID]int64)

	sim := simulation.New(map[string]simulation.ServiceFunc{
		"swarm": func(ctx *adapters.ServiceContext, bucket *sync.Map) (s node.Service, cleanup func(), err error) {
			config := api.NewConfig()
			config.Port = ""

			dir, err := ioutil.TempDir("", "swap-network-test-node")
			if err != nil {
				return nil, nil, err
			}

			cleanup = func() {
				err := os.RemoveAll(dir)
				if err != nil {
					log.Error("cleaning up swarm temp dir", "err", err)
				}
			}

			config.Path = dir

			privkey, err := crypto.GenerateKey()
			if err != nil {
				return nil, cleanup, err
			}

			config.Init(privkey)
			//enable swap
			config.SwapEnabled = true

			swarm, err := NewSwarm(config, nil)
			if err != nil {
				return nil, cleanup, err
			}
			bucket.Store(bucketKeySwarm, swarm)
			log.Info("new swarm", "bzzKey", config.BzzKey, "baseAddr", fmt.Sprintf("%x", swarm.bzz.BaseAddr()))
			return swarm, cleanup, nil
		},
	})
	defer sim.Close()

	ctx := context.Background()
	files := make([]file, 0)

	//upload a snapshot
	err := sim.UploadSnapshot(fmt.Sprintf("network/stream/testing/snapshot_%d.json", nodeCount))
	if err != nil {
		t.Fatal(err)
	}

	//NOTE: maxFileSize is 4 kB, this in order to provide faster tests
	//it would be interesting to run these tests with bigger files
	//(to see how drop limits are affected etc.)
	const maxFileSize = 1024 * 4 //1024 bytes * 4 = 4kB
	const minfileSize = 1024

	//pseudo random algo to define if a node will upload or not
	//if a bit is 0, do not upload
	pseudoRandomNum := rand.Int63()
	pseudoRandomBitMask := strconv.FormatInt(pseudoRandomNum, 2)

	result := sim.Run(ctx, func(ctx context.Context, sim *simulation.Simulation) error {
		if _, err := sim.WaitTillHealthy(ctx, 2); err != nil {
			return err
		}

		nodeIDs := sim.UpNodeIDs()
		rand.Shuffle(len(nodeIDs), func(i, j int) {
			nodeIDs[i], nodeIDs[j] = nodeIDs[j], nodeIDs[i]
		})
		for i, id := range nodeIDs {
			//if the position in random num is 0, don't upload
			if string(pseudoRandomBitMask[i]) != "0" {
				size := rand.Intn(maxFileSize-minfileSize) + minfileSize
				key, data, err := uploadRandomFileSize(sim.Service("swarm", id).(*Swarm), size)
				if err != nil {
					return err
				}
				log.Trace("file uploaded", "node", id, "key", key.String())
				files = append(files, file{
					addr:   key,
					data:   data,
					nodeID: id,
				})
			}
		}

		// File retrieval check is repeated until all uploaded files are retrieved from all nodes
		// or until the timeout is reached.
		//we use a special retrieve function for swap which is optimized for parallel requests
		//but does not leave many cascaded requests floating around
		for retrieveForSwap(sim, files) != 0 {
		}

		for _, node := range sim.NodeIDs() {
			item, ok := sim.NodeItem(node, bucketKeySwarm)
			if !ok {
				return errors.New("no swarm")
			}
			swarm := item.(*Swarm)

			subBalances := make(map[enode.ID]int64)

			for _, n := range sim.NodeIDs() {
				if node == n {
					continue
				}
				balance, err := swarm.swap.GetPeerBalance(n)
				if err == nil {
					subBalances[n] = balance
					log.Debug(fmt.Sprintf("Balance of node %s to node %s: %d", node.TerminalString(), n.TerminalString(), balance))
				} else {
					log.Debug(fmt.Sprintf("Node %s has no balance with node %s", node.TerminalString(), n.TerminalString()))
				}
			}
			balancesMap[node] = subBalances
		}
		return nil
	})

	if result.Error != nil {
		t.Fatal(result.Error)
	}
	if *printStats {
		for k, v := range balancesMap {
			fmt.Println(fmt.Sprintf("node %s balances:", k.TerminalString()))
			for kk, vv := range v {
				fmt.Println(fmt.Sprintf(".........with node %s: balance %d", kk.TerminalString(), vv))
			}
		}
	}

	/*
		Assuming that in this case, balances should be symmetric too	I
	*/

	errorsFound := false
	for k, mapForK := range balancesMap {
		for n, balanceKwithN := range mapForK {
			mapForSubK := balancesMap[n]
			log.Trace(fmt.Sprintf("balance of %s with %s: %d", k.TerminalString(), n.TerminalString(), balanceKwithN))
			log.Trace(fmt.Sprintf("balance of %s with %s: %d", n.TerminalString(), k.TerminalString(), mapForSubK[k]))
			if balanceKwithN+mapForSubK[k] != 0 && balanceKwithN != 0 {
				log.Error(fmt.Sprintf("Expected balances to be a+b = 0 AND balance(a) != 0, but they are not, balance k with n: %d, balance n with k: %d", balanceKwithN, mapForSubK[k]))
				errorsFound = true
			}
		}
	}

	if errorsFound {
		t.Fatal("Expected balances to be symmetrical, but they were not")
	}
	log.Debug("test terminated")
}

// uploadRandomFileSize, uploads a file of random size to the swarm instance
// using the api.Put method.
func uploadRandomFileSize(swarm *Swarm, size int) (storage.Address, string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, "", err
	}
	// uniqueness is very certain.
	data := fmt.Sprintf("test content %s %x", time.Now().Round(0), b)
	ctx := context.TODO()
	k, wait, err := swarm.api.Put(ctx, data, "text/plain", false)
	if err != nil {
		return nil, "", err
	}
	if wait != nil {
		err = wait(ctx)
	}
	return k, data, err
}

// retrieveForSwap is a special retrieve function for swap tests which is slightly
// optimized for parallel request but does not leave many cascaded requests floating around
func retrieveForSwap(
	sim *simulation.Simulation,
	files []file,
) (missing uint64) {
	rand.Shuffle(len(files), func(i, j int) {
		files[i], files[j] = files[j], files[i]
	})

	nodeIDs := sim.UpNodeIDs()

	totalFoundCount := uint64(0)
	totalCheckCount := uint64(len(nodeIDs) * len(files))

	for _, id := range nodeIDs {

		swarm := sim.Service("swarm", id).(*Swarm)
		for _, f := range files {
			log.Debug("api get: check file", "node", id.String(), "key", f.addr.String(), "total files found", totalFoundCount)

			r, _, _, _, err := swarm.api.Get(context.TODO(), api.NOOPDecrypt, f.addr, "/")
			if err != nil {
				log.Error("api get: node %s, key %s, kademlia %s: %v", id, f.addr, swarm.bzz.Hive, err)
				return
			}
			d, err := ioutil.ReadAll(r)
			if err != nil {
				log.Error("api get: read response: node %s, key %s: kademlia %s: %v", id, f.addr, swarm.bzz.Hive, err)
				return
			}
			data := string(d)
			if data != f.data {
				log.Error("file contend missmatch: node %s, key %s, expected %q, got %q", id, f.addr, f.data, data)
				return
			}
			log.Info("api get: file found", "node", id.String(), "key", f.addr.String(), "content", data, "files found", totalFoundCount)

			totalFoundCount++

			log.Debug("status", "totalCheckCount", totalCheckCount, "totalFoundCount", totalFoundCount)
		}
	}

	log.Info("check stats", "total check count", totalCheckCount, "total files found", totalFoundCount, "missing", missing)

	return totalCheckCount - totalFoundCount
}

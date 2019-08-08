// Copyright 2019 The go-ethereum Authors
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

package snapshot

import (
	"fmt"
	"io"
	"os"
	"sort"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
)

// diffLayer represents a collection of modifications made to a state snapshot
// after running a block on top. It contains one sorted list for the account trie
// and one-one list for each storage tries.
//
// The goal of a diff layer is to act as a journal, tracking recent modifications
// made to the state, that have not yet graduated into a semi-immutable state.
type diffLayer struct {
	parent snapshot // Parent snapshot modified by this one, never nil
	memory uint64   // Approximate guess as to how much memory we use

	number uint64      // Block number to which this snapshot diff belongs to
	root   common.Hash // Root hash to which this snapshot diff belongs to

	accountList   []common.Hash                          // List of account for iteration, might not be sorted yet (lazy)
	accountSorted bool                                   // Flag whether the account list has alreayd been sorted or not
	accountData   map[common.Hash][]byte                 // Keyed accounts for direct retrival (nil means deleted)
	storageList   map[common.Hash][]common.Hash          // List of storage slots for iterated retrievals, one per account
	storageSorted map[common.Hash]bool                   // Flag whether the storage slot list has alreayd been sorted or not
	storageData   map[common.Hash]map[common.Hash][]byte // Keyed storage slots for direct retrival. one per account (nil means deleted)

	lock sync.RWMutex
}

// newDiffLayer creates a new diff on top of an existing snapshot, whether that's a low
// level persistent database or a hierarchical diff already.
func newDiffLayer(parent snapshot, root common.Hash, accounts map[common.Hash][]byte, storage map[common.Hash]map[common.Hash][]byte) *diffLayer {
	// Create the new layer with some pre-allocated data segments
	parentNumber, _ := parent.Info()

	dl := &diffLayer{
		parent:      parent,
		number:      parentNumber + 1,
		root:        root,
		accountData: accounts,
		storageData: storage,
	}
	// Fill the account hashes and sort them for the iterator
	accountList := make([]common.Hash, 0, len(accounts))
	for hash, data := range accounts {
		accountList = append(accountList, hash)
		dl.memory += uint64(len(data))
	}
	sort.Sort(hashes(accountList))
	dl.accountList = accountList
	dl.accountSorted = true

	dl.memory += uint64(len(dl.accountList) * common.HashLength)

	// Fill the storage hashes and sort them for the iterator
	dl.storageList = make(map[common.Hash][]common.Hash, len(storage))
	dl.storageSorted = make(map[common.Hash]bool, len(storage))

	for accountHash, slots := range storage {
		// If the slots are nil, sanity check that it's a deleted account
		if slots == nil {
			// Ensure that the account was just marked as deleted
			if account, ok := accounts[accountHash]; account != nil || !ok {
				panic(fmt.Sprintf("storage in %#x nil, but account conflicts (%#x, exists: %v)", accountHash, account, ok))
			}
			// Everything ok, store the deletion mark and continue
			dl.storageList[accountHash] = nil
			continue
		}
		// Storage slots are not nil so entire contract was not deleted, ensure the
		// account was just updated.
		if account, ok := accounts[accountHash]; account == nil || !ok {
			log.Error(fmt.Sprintf("storage in %#x exists, but account nil (exists: %v)", accountHash, ok))
			//panic(fmt.Sprintf("storage in %#x exists, but account nil (exists: %v)", accountHash, ok))
		}
		// Fill the storage hashes for this account and sort them for the iterator
		storageList := make([]common.Hash, 0, len(slots))
		for storageHash, data := range slots {
			storageList = append(storageList, storageHash)
			dl.memory += uint64(len(data))
		}
		sort.Sort(hashes(storageList))
		dl.storageList[accountHash] = storageList
		dl.storageSorted[accountHash] = true

		dl.memory += uint64(len(storageList) * common.HashLength)
	}
	dl.memory += uint64(len(dl.storageList) * common.HashLength)

	return dl
}

// loadDiffLayer reads the next sections of a snapshot journal, reconstructing a new
// diff and verifying that it can be linked to the requested parent.
func loadDiffLayer(parent snapshot, r io.Reader) (snapshot, error) {
	// Read the next diff journal entry
	var (
		number   uint64
		root     common.Hash
		accounts = make(map[common.Hash][]byte)
		storage  = make(map[common.Hash]map[common.Hash][]byte)
	)
	if err := rlp.Decode(r, &number); err != nil {
		// The first read may fail with EOF, marking the end of the journal
		if err == io.EOF {
			return parent, nil
		}
		return nil, err
	}
	if err := rlp.Decode(r, &root); err != nil {
		return nil, err
	}
	if err := rlp.Decode(r, &accounts); err != nil {
		return nil, err
	}
	if err := rlp.Decode(r, &storage); err != nil {
		return nil, err
	}
	// Validate the block number to avoid state corruption
	if parentNumber, _ := parent.Info(); number != parentNumber+1 {
		return nil, fmt.Errorf("snapshot chain broken: block #%dl after #%dl", number, parentNumber)
	}
	return loadDiffLayer(newDiffLayer(parent, root, accounts, storage), r)
}

// Info returns the block number and root hash for which this snapshot was made.
func (dl *diffLayer) Info() (uint64, common.Hash) {
	return dl.number, dl.root
}

// Account directly retrieves the account associated with a particular hash in
// the snapshot slim data format.
func (dl *diffLayer) Account(hash common.Hash) *Account {
	data := dl.AccountRLP(hash)
	if len(data) == 0 { // nil if deleted, []byte{} if never existed
		return nil
	}
	account := new(Account)
	if err := rlp.DecodeBytes(data, account); err != nil {
		panic(err)
	}
	return account
}

// AccountRLP directly retrieves the account RLP associated with a particular
// hash in the snapshot slim data format.
func (dl *diffLayer) AccountRLP(hash common.Hash) []byte {
	dl.lock.RLock()
	defer dl.lock.RUnlock()

	// If the account is known locally, return it. Note, a nil account means it was
	// deleted, and is a different notion than an unknown account!
	if data, ok := dl.accountData[hash]; ok {
		return data
	}
	// Account unknown to this diff, resolve from parent
	return dl.parent.AccountRLP(hash)
}

// Storage directly retrieves the storage data associated with a particular hash,
// within a particular account. If the slot is unknown to this diff, it's parent
// is consulted.
func (dl *diffLayer) Storage(accountHash, storageHash common.Hash) []byte {
	dl.lock.RLock()
	defer dl.lock.RUnlock()

	// If the account is known locally, try to resolve the slot locally. Note, a nil
	// account means it was deleted, and is a different notion than an unknown account!
	if storage, ok := dl.storageData[accountHash]; ok {
		if storage == nil {
			return nil
		}
		if data, ok := storage[storageHash]; ok {
			return data
		}
	}
	// Account - or slot within - unknown to this diff, resolve from parent
	return dl.parent.Storage(accountHash, storageHash)
}

// Update creates a new layer on top of the existing snapshot diff tree with
// the specified data items.
func (dl *diffLayer) Update(blockRoot common.Hash, accounts map[common.Hash][]byte, storage map[common.Hash]map[common.Hash][]byte) *diffLayer {
	return newDiffLayer(dl, blockRoot, accounts, storage)
}

// Cap traverses downwards the diff tree until the number of allowed layers are
// crossed. All diffs beyond the permitted number are flattened downwards. If
// the layer limit is reached, memory cap is also enforced (but not before). The
// block numbers for the disk layer and first diff layer are returned for GC.
func (dl *diffLayer) Cap(layers int, memory uint64) (uint64, uint64) {
	// Dive until we run out of layers or reach the persistent database
	if layers > 2 {
		// If we still have diff layers below, recurse
		if parent, ok := dl.parent.(*diffLayer); ok {
			return parent.Cap(layers-1, memory)
		}
		// Diff stack too shallow, return block numbers without modifications
		return dl.parent.(*diskLayer).number, dl.number
	}
	// We're out of layers, flatten anything below, stopping if it's the disk or if
	// the memory limit is not yet exceeded.
	switch parent := dl.parent.(type) {
	case *diskLayer:
		return parent.number, dl.number
	case *diffLayer:
		dl.lock.Lock()
		defer dl.lock.Unlock()

		dl.parent = parent.flatten()
		if dl.parent.(*diffLayer).memory < memory {
			diskNumber, _ := parent.parent.Info()
			return diskNumber, parent.number
		}
	default:
		panic(fmt.Sprintf("unknown data layer: %T", parent))
	}
	// If the bottommost layer is larger than our memory cap, persist to disk
	var (
		parent = dl.parent.(*diffLayer)
		base   = parent.parent.(*diskLayer)
		batch  = base.db.NewBatch()
	)
	// Start by temporarilly deleting the current snapshot block marker. This
	// ensures that in the case of a crash, the entire snapshot is invalidated.
	rawdb.DeleteSnapshotBlock(batch)

	// Push all the accounts into the database
	for hash, data := range parent.accountData {
		rawdb.WriteAccountSnapshot(batch, hash, data)
		if batch.ValueSize() > ethdb.IdealBatchSize {
			if err := batch.Write(); err != nil {
				log.Crit("Failed to write account snapshot", "err", err)
			}
			batch.Reset()
		}
	}
	// Push all the storage slots into the database
	for accountHash, storage := range parent.storageData {
		for storageHash, data := range storage {
			rawdb.WriteStorageSnapshot(batch, accountHash, storageHash, data)
		}
		if batch.ValueSize() > ethdb.IdealBatchSize {
			if err := batch.Write(); err != nil {
				log.Crit("Failed to write storage snapshot", "err", err)
			}
			batch.Reset()
		}
	}
	// Update the snapshot block marker and write any remainder data
	base.number, base.root = parent.number, parent.root

	rawdb.WriteSnapshotBlock(batch, base.number, base.root)
	if err := batch.Write(); err != nil {
		log.Crit("Failed to write leftover snapshot", "err", err)
	}
	dl.parent = base

	return base.number, dl.number
}

// flatten pushes all data from this point downwards, flattening everything into
// a single diff at the bottom. Since usually the lowermost diff is the largest,
// the flattening bulds up from there in reverse.
func (dl *diffLayer) flatten() snapshot {
	// If the parent is not diff, we're the first in line, return unmodified
	parent, ok := dl.parent.(*diffLayer)
	if !ok {
		return dl
	}
	// Parent is a diff, flatten it first (note, apart from weird corned cases,
	// flatten will realistically only ever merge 1 layer, so there's no need to
	// be smarter about grouping flattens together).
	parent = parent.flatten().(*diffLayer)

	// Overwrite all the updated accounts blindly, merge the sorted list
	for hash, data := range dl.accountData {
		parent.accountData[hash] = data
	}
	parent.accountList = append(parent.accountList, dl.accountList...) // TODO(karalabe): dedup!!
	parent.accountSorted = false

	// Overwrite all the updates storage slots (individually)
	for accountHash, storage := range dl.storageData {
		// If storage didn't exist (or was deleted) in the parent; or if the storage
		// was freshly deleted in the child, overwrite blindly
		if parent.storageData[accountHash] == nil || storage == nil {
			parent.storageList[accountHash] = dl.storageList[accountHash]
			parent.storageData[accountHash] = storage
			continue
		}
		// Storage exists in both parent and child, merge the slots
		comboData := parent.storageData[accountHash]
		for storageHash, data := range storage {
			comboData[storageHash] = data
		}
		parent.storageData[accountHash] = comboData
		parent.storageList[accountHash] = append(parent.storageList[accountHash], dl.storageList[accountHash]...) // TODO(karalabe): dedup!!
		parent.storageSorted[accountHash] = false
	}
	// Return the combo parent
	parent.number = dl.number
	parent.root = dl.root
	parent.memory += dl.memory
	return parent
}

// Journal commits an entire diff hierarchy to disk into a single journal file.
// This is meant to be used during shutdown to persist the snapshot without
// flattening everything down (bad for reorgs).
func (dl *diffLayer) Journal() error {
	dl.lock.RLock()
	defer dl.lock.RUnlock()

	writer, err := dl.journal()
	if err != nil {
		return err
	}
	writer.Close()
	return nil
}

// journal is the internal version of Journal that also returns the journal file
// so subsequent layers know where to write to.
func (dl *diffLayer) journal() (io.WriteCloser, error) {
	// If we've reached the bottom, open the journal
	var writer io.WriteCloser
	if parent, ok := dl.parent.(*diskLayer); ok {
		file, err := os.Create(parent.journal)
		if err != nil {
			return nil, err
		}
		writer = file
	}
	// If we haven't reached the bottom yet, journal the parent first
	if writer == nil {
		file, err := dl.parent.(*diffLayer).journal()
		if err != nil {
			return nil, err
		}
		writer = file
	}
	// Everything below was journalled, persist this layer too
	if err := rlp.Encode(writer, dl.number); err != nil {
		writer.Close()
		return nil, err
	}
	if err := rlp.Encode(writer, dl.root); err != nil {
		writer.Close()
		return nil, err
	}
	if err := rlp.Encode(writer, dl.accountData); err != nil {
		writer.Close()
		return nil, err
	}
	if err := rlp.Encode(writer, dl.storageData); err != nil {
		writer.Close()
		return nil, err
	}
	return writer, nil
}

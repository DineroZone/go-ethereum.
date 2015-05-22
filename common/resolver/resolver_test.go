package resolver

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type testBackend struct {
	// contracts mock
	contracts map[string](map[string]string)
}

var (
	text     = "test"
	codehash = common.StringToHash("1234")
	hash     = common.BytesToHash(crypto.Sha3([]byte(text)))
	url      = "bzz://bzzhash/my/path/contr.act"
)

func NewTestBackend() *testBackend {
	HashRegContractAddress = common.BigToAddress(common.Big0).Hex() //[2:]
	UrlHintContractAddress = common.BigToAddress(common.Big1).Hex() //[2:]
	self := &testBackend{}
	self.contracts = make(map[string](map[string]string))

	self.contracts[HashRegContractAddress[2:]] = make(map[string]string)
	key := storageAddress(storageMapping(storageIdx2Addr(1), codehash[:]))
	self.contracts[HashRegContractAddress[2:]][key] = hash.Hex()

	self.contracts[UrlHintContractAddress[2:]] = make(map[string]string)
	mapaddr := storageMapping(storageIdx2Addr(1), hash[:])

	key = storageAddress(storageFixedArray(mapaddr, storageIdx2Addr(0)))
	self.contracts[UrlHintContractAddress[2:]][key] = common.ToHex([]byte(url))
	key = storageAddress(storageFixedArray(mapaddr, storageIdx2Addr(1)))
	self.contracts[UrlHintContractAddress[2:]][key] = "0x0"
	return self
}

func (self *testBackend) StorageAt(ca, sa string) (res string) {
	c := self.contracts[ca]
	if c == nil {
		return
	}
	res = c[sa]
	return
}

func (self *testBackend) Transact(fromStr, toStr, nonceStr, valueStr, gasStr, gasPriceStr, codeStr string) (string, error) {
	return "", nil
}

func (self *testBackend) Call(fromStr, toStr, valueStr, gasStr, gasPriceStr, codeStr string) (string, string, error) {
	return "", "", nil
}

func TestCreateContractsExists(t *testing.T) {
	b := NewTestBackend()
	res := New(b)
	_, _, err := res.CreateContracts(common.BigToAddress(common.Big0))
	exp := "HashReg already exists at 0x0000000000000000000000000000000000000000"
	if err == nil || err.Error() != exp {
		t.Errorf("expected error '%s', got '%v'", exp, err)
	}
}

func TestKeyToContentHash(t *testing.T) {
	b := NewTestBackend()
	res := New(b)

	got, err := res.KeyToContentHash(codehash)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	} else {
		if got != hash {
			t.Errorf("incorrect result, expected '%v', got '%v'", hash.Hex(), got.Hex())
		}
	}
}

func TestContentHashToUrl(t *testing.T) {
	b := NewTestBackend()
	res := New(b)
	got, err := res.ContentHashToUrl(hash)
	if err != nil {
		t.Errorf("expected 	 error, got %v", err)
	} else {
		if got != url {
			t.Errorf("incorrect result, expected '%v', got '%s'", url, got)
		}
	}
}

func TestKeyToUrl(t *testing.T) {
	b := NewTestBackend()
	res := New(b)
	got, _, err := res.KeyToUrl(codehash)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	} else {
		if got != url {
			t.Errorf("incorrect result, expected \n'%s', got \n'%s'", url, got)
		}
	}
}

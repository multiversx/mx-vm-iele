package main

import (
	"bytes"

	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"

	worldhook "github.com/ElrondNetwork/elrond-vm-util/mock-hook-blockchain"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	ij "github.com/ElrondNetwork/elrond-vm-util/test-util/vmtestjson"
)

func convertAccount(testAcct *ij.Account) *worldhook.Account {
	storage := make(map[string][]byte)
	for _, stkvp := range testAcct.Storage {
		if stkvp.Value == nil {
			panic("why?")
		}
		key := normalizeStorageKey(string(stkvp.Key))
		storage[key] = stkvp.Value
	}

	return &worldhook.Account{
		Exists:  true,
		Address: testAcct.Address,
		Nonce:   testAcct.Nonce.Uint64(),
		Balance: big.NewInt(0).Set(testAcct.Balance),
		Storage: storage,
		Code:    []byte(testAcct.Code),
	}
}

func convertLogToTestFormat(outputLog *vmi.LogEntry) *ij.LogEntry {
	testLog := ij.LogEntry{
		Address: outputLog.Address,
		Topics:  outputLog.Topics,
		Data:    outputLog.Data,
	}
	return &testLog
}

var zero = big.NewInt(0)

func zeroIfNil(i *big.Int) *big.Int {
	if i == nil {
		return zero
	}
	return i
}

// tool to modify tests
// use with caution
func saveModifiedTest(toPath string, top []*ij.Test) {
	resultJSON := ij.ToJSONString(top)

	err := os.MkdirAll(filepath.Dir(toPath), os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(toPath, []byte(resultJSON), 0644)
	if err != nil {
		panic(err)
	}
}

func normalizeStorageKey(key string) string {
	return string(big.NewInt(0).SetBytes([]byte(key)).Bytes())
	// trimmed := strings.TrimLeft(key, "0")
	// if len(trimmed)%2 == 1 {
	// 	trimmed = "0" + trimmed
	// }
	// return trimmed
}

func storageValuesEqual(value1, value2 []byte) bool {
	if bytes.Equal(value1, value2) {
		return true
	}
	bi1 := big.NewInt(0).SetBytes(value1)
	bi2 := big.NewInt(0).SetBytes(value2)
	return bi1.Cmp(bi2) == 0
}

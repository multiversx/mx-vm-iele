package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"

	worldhook "github.com/ElrondNetwork/elrond-vm-util/mock-hook-blockchain"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	ij "github.com/ElrondNetwork/elrond-vm-util/test-util/ielejson"
)

// for nicer error messages
func resultAsString(result []*big.Int) string {
	str := "["
	for i, res := range result {
		str += fmt.Sprintf("0x%x", res)
		if i < len(result)-1 {
			str += ", "
		}
	}
	return str + "]"
}

func convertAccount(testAcct *ij.Account) *worldhook.Account {
	storage := make(map[string][]byte)
	for _, stkvp := range testAcct.Storage {
		if stkvp.Value == nil {
			panic("why?")
		}
		key := string(stkvp.Key.Bytes())
		storage[key] = stkvp.Value.Bytes()
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

func convertBlockHeader(testBlh *ij.BlockHeader) *vmi.SCCallHeader {
	return &vmi.SCCallHeader{
		Beneficiary: testBlh.Beneficiary,
		Number:      testBlh.Number,
		GasLimit:    testBlh.GasLimit,
		Timestamp:   testBlh.UnixTimestamp,
	}
	// return &vmi.SCCallHeader{
	// 	Beneficiary: big.NewInt(0),
	// 	Number:      big.NewInt(0),
	// 	GasLimit:    big.NewInt(0),
	// 	Timestamp:   big.NewInt(0),
	// }
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

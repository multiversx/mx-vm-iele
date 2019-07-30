package main

import (
	"fmt"
	"path"
)

func main() {
	//debugTest("tests/iele/danse/ill-formed/illFormed2.iele.json")
	//debugTest("tests/iele-v2/albe/ERC20/approve_Caller-Positive.iele.json")
	//debugVMTest("tests/VMTests/vmPerformance/loop-mul/loop-mul.iele.json", false)
	debugVMTest("tests/VMTests/vmPerformance/loop-exp-nop-1M/loop-exp-nop-1M.iele.json", false)
}

// .build/vm/iele-test-vm tests/iele/danse/forwarder/create.iele.json 10000

func debugIeleTest(testFile string, tracePretty bool) {
	err := runTest(path.Join(ieleTestRoot, testFile), gasModeNormal, tracePretty, nil)
	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}

func debugVMTest(testFile string, tracePretty bool) {
	err := runTest(path.Join(ieleTestRoot, testFile), gasModeVMTests, tracePretty, nil)
	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}

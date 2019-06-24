package main

import (
	"fmt"
	"path"
)

func main() {
	debugTest("tests/iele/danse/ill-formed/illFormed2.iele.json")
	//debugTest("tests/iele-v2/albe/ERC20/approve_Caller-Positive.iele.json")
}

// .build/vm/iele-test-vm tests/iele/danse/forwarder/create.iele.json 10000

func debugTest(testFile string) {
	err := runTest(path.Join(ieleTestRoot, testFile), gasModeNormal, true)
	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		panic(err)
	}
}

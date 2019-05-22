package main

import (
	"fmt"
	"path"

	interpreter "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestinginterpreter"
)

func main() {
	if false {
		interpreter.StartTraceReferenceComparer("ocaml_or3_1.log")
	}

	// debugTest("tests/iele/danse/ill-formed/illFormed2.iele.json")
	// debugTest("tests/iele/danse/factorial/factorial_positive.iele.json")
	debugTest("tests/iele/albe/unit/exceptions.iele.json")

}

// .build/vm/iele-test-vm tests/iele/danse/forwarder/create.iele.json 10000

func debugTest(testFile string) {
	err := runTest(path.Join(ieleTestRoot, testFile), gasModeNormal, false, false)
	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		panic(err)
	}
}

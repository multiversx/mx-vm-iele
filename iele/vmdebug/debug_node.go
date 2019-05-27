package main

import (
	"fmt"
	"path"
	"path/filepath"

	eiele "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/endpoint"
	oiele "github.com/ElrondNetwork/elrond-vm/iele/original/node/endpoint"
	eptest "github.com/ElrondNetwork/elrond-vm/iele/test-util/endpointtest"

	einterpreter "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestinginterpreter"
	ointerpreter "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestinginterpreter"
)

func main() {

	//debugTest("tests/iele/danse/factorial/factorial_positive.iele.json")
	//debugTest("tests/iele/albe/factorial/factorial_positive.iele.json")
	//debugTest("tests/iele/albe/ERC20/approve_Caller-Positive.iele.json")
	//debugTest("tests/iele/albe/unit/blockhash.iele.json")
	//debugTest("tests/iele/albe/unit/exceptions.iele.json")
	//debugTest("tests/iele-v1/albe/forwarder/create.iele.json")
	debugElrondTest("tests/iele-v2/danse/forwarder/copycreate.iele.json", false)

}

func debugElrondTest(testFile string, tracePretty bool) {
	if tracePretty {
		eiele.InterpreterOptions = &einterpreter.ExecuteOptions{
			TracePretty: true,
			TraceKPrint: false,
			Verbose:     false,
			MaxSteps:    0,
		}
	}
	err := eptest.RunJSONTest(path.Join(elrondTestRoot, testFile), eiele.ElrondIeleVM)
	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		fmt.Printf("ERROR: %s", err.Error())
	}
}

func debugOriginalTest(testFile string, tracePretty bool) {
	if tracePretty {
		oiele.InterpreterOptions = &ointerpreter.ExecuteOptions{
			TracePretty: true,
			TraceKPrint: false,
			Verbose:     false,
			MaxSteps:    0,
		}
	}
	err := eptest.RunJSONTest(filepath.Join(originalTestRoot, testFile), oiele.OriginalIeleVM)
	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		fmt.Printf("ERROR: %s", err.Error())
	}
}

package main

import (
	"fmt"
	"path"
	"path/filepath"

	eptest "github.com/ElrondNetwork/elrond-vm/iele/test-util/endpointtest"
)

func main() {

	//debugTest("tests/iele/danse/factorial/factorial_positive.iele.json")
	//debugTest("tests/iele/albe/factorial/factorial_positive.iele.json")
	//debugTest("tests/iele/albe/ERC20/approve_Caller-Positive.iele.json")
	//debugTest("tests/iele/albe/unit/blockhash.iele.json")
	//debugTest("tests/iele-v1/albe/forwarder/create.iele.json")
	//debugElrondTest("tests/iele-v2/danse/forwarder/copycreate.iele.json", false)

	// debugOriginalTest("tests/iele/albe/unit/exceptions.iele.json", false)
	// debugOriginalTest("tests/iele/danse/unit/exceptions.iele.json", false)
	// debugElrondTest("tests/iele-v2/albe/unit/exceptions.iele.json", false)
	// debugElrondTest("tests/iele-v2/danse/unit/exceptions.iele.json", false)

	debugElrondTest("agar/create.iele.json", false)
	debugElrondTest("agar/topUp1.iele.json", false)
	debugElrondTest("agar/balanceOf.iele.json", false)
	debugElrondTest("agar/withdraw_Ok.iele.json", false)
	debugElrondTest("agar/withdraw_TooMuch.iele.json", false)
	debugElrondTest("agar/withdrawAll_Ok.iele.json", false)
	debugElrondTest("agar/addPlayerToGame.iele.json", false)
	debugElrondTest("agar/rewardWinner.iele.json", false)
	debugElrondTest("agar/endGame.iele.json", false)
	debugElrondTest("agar/exceptions.iele.json", false)
}

func debugElrondTest(testFile string, tracePretty bool) {
	// if tracePretty {
	// 	eiele.InterpreterOptions = &einterpreter.ExecuteOptions{
	// 		TracePretty: true,
	// 		TraceKPrint: false,
	// 		Verbose:     false,
	// 		MaxSteps:    0,
	// 	}
	// }
	err := eptest.RunJSONTest(path.Join(elrondTestRoot, testFile), elrondIeleProvider{}, world)
	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}

func debugOriginalTest(testFile string, tracePretty bool) {
	// if tracePretty {
	// 	oiele.InterpreterOptions = &ointerpreter.ExecuteOptions{
	// 		TracePretty: true,
	// 		TraceKPrint: false,
	// 		Verbose:     false,
	// 		MaxSteps:    0,
	// 	}
	// }
	err := eptest.RunJSONTest(filepath.Join(originalTestRoot, testFile), originalIeleProvider{}, world)
	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}

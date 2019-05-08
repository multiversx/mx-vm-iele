package main

import (
	"fmt"
	interpreter "kgoieletesting/iele-testing-kompiled/ieletestinginterpreter"
	"path"
)

func main() {
	if false {
		interpreter.StartTraceReferenceComparer("ocaml_or3_1.log")
	}

	// debugTest("tests/iele/danse/ill-formed/illFormed2.iele.json")
	// debugTest("tests/iele/danse/factorial/factorial_positive.iele.json")
	// debugTest("tests/iele/danse/forwarder/create.iele.json")

	debugTest("tests/iele/danse/ERC20/transferFrom_AllDistinct-BalanceEqAllowance.iele.json")

}

func debugTest(testFile string) {
	err := runTest(path.Join(ieleTestRoot, testFile), gasModeNormal, false, false)
	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		panic(err)
	}
}

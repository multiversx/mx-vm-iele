package main

import (
	"fmt"
	"path"
	"path/filepath"

	eptest "github.com/ElrondNetwork/elrond-vm/iele/test-util/endpointtest"
)

func main() {

	//debugElrondTest("tests/iele-v2/danse/forwarder/copycreate.iele.json", false)

	//debugElrondTest("tests/iele-v2/albe/ill-formed/illFormed2.iele.json", false)
	//debugOriginalTest("tests/iele/albe/ill-formed/illFormed2.iele.json", false)

	debugElrondTest("tests/iele-v3/danse/ERC20/transferFrom_FromEqTo-BalanceEqAllowance.iele.json", false)

	//debugElrondTest("adder/adder.iele.json", false)
	//debugAgar()
}

func debugAgar() {
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
	err := eptest.RunJSONTest(
		path.Join(elrondTestRoot, testFile),
		&elrondIeleProvider{tracePretty: tracePretty},
		world)
	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}

func debugOriginalTest(testFile string, tracePretty bool) {
	err := eptest.RunJSONTest(
		filepath.Join(originalTestRoot, testFile),
		&originalIeleProvider{tracePretty: tracePretty},
		world)
	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}

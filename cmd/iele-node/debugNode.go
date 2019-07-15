package main

import (
	"fmt"
	"path/filepath"

	controller "github.com/ElrondNetwork/elrond-vm/iele/test-util/testcontroller"
)

func main() {

	//debugElrondTest("tests/iele-v2/danse/forwarder/copycreate.iele.json", false)

	//debugElrondTest("tests/iele-v3/danse/ill-formed/illFormed2.iele.json", false)
	//debugOriginalTest("tests/iele/danse/ill-formed/illFormed2.iele.json", false)
	//debugOriginalTest("tests/iele/albe/ill-formed/illFormedX14.iele.json", false)
	debugOriginalTest("tests/iele/albe/ERC20/create.iele.json", false)

	//debugElrondTest("auction-solidity/create.iele.json", false)
	debugElrondTest("agar/addPlayerToGame.iele.json", false)

	//debugElrondTest("adder/adder.iele.json", false)
	//debugElrondTest("tests/iele-v3/danse/ERC20/allowance_CallerCaller.iele.json", false)
	//ebugElrondTest("agar_v1/endGame.iele.json", true)
	//debugAgarV2()

	//debugIllFormedX()
}

func debugAgarV2() {
	debugElrondTest("agar_v2/create.iele.json", false)
	debugElrondTest("agar_v2/topUp1.iele.json", false)
	debugElrondTest("agar_v2/balanceOf.iele.json", false)
	debugElrondTest("agar_v2/withdraw_Ok.iele.json", false)
	debugElrondTest("agar_v2/withdraw_TooMuch.iele.json", false)
	debugElrondTest("agar_v2/withdrawAll_Ok.iele.json", false)
	debugElrondTest("agar_v2/addPlayerToGame.iele.json", false)
	debugElrondTest("agar_v2/joinGame.iele.json", false)
	debugElrondTest("agar_v2/rewardWinner.iele.json", false)
	debugElrondTest("agar_v2/endGame.iele.json", false)
	debugElrondTest("agar_v2/endGame_leftover.iele.json", false)
	debugElrondTest("agar_v2/exceptions.iele.json", false)
}

func debugElrondTest(testFile string, tracePretty bool) {
	err := controller.RunSingleIeleTest(
		filepath.Join(elrondTestRoot, testFile),
		newElrondIeleTestExecutor(tracePretty))

	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}

func debugOriginalTest(testFile string, tracePretty bool) {
	err := controller.RunSingleIeleTest(
		filepath.Join(originalTestRoot, testFile),
		newOriginalIeleTestExecutor(tracePretty))

	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}

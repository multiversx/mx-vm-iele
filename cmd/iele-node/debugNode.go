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
	//debugOriginalTest("tests/iele/albe/ERC20/create.iele.json", false)

	//debugElrondTest("auction-solidity/create.iele.json", false)
	//debugElrondTestnetTest("agar_v2/withdraw_TooMuch.iele.json", true)

	//debugElrondTest("agar_v2/endGame_leftover.iele.json", false)

	//debugOriginalTest("perf-solidity/LoopExpNop1M.json", false)
	//debugOriginalTest("perf-solidity/LoopDivAdd10M.json", false)

	//debugElrondTest("agar_v1/addPlayerToGame.iele.json", false)
	//debugElrondTest("tests/iele-v3/danse/ERC20/allowance_CallerCaller.iele.json", false) // trace_good
	//debugElrondTest("tests/iele-v3/danse/ERC20/approve_Caller-Zero.iele.json", true)

	//debugAgarMin()

	debugElrondTest("agar_v4/topUp_ok.iele.json", true)
	//debugElrondTest("tests/iele-v3/danse/unit/exceptions.iele.json", false)

}

func debugAgarMin() {
	debugElrondTest("agar_min_v1/joinGame.iele.json", false)

	debugElrondTest("agar_min_v1/create.iele.json", false)
	debugElrondTest("agar_min_v1/topUp_ok.iele.json", false)
	debugElrondTest("agar_min_v1/topUp_outOfFunds.iele.json", false)
	debugElrondTest("agar_min_v1/topUp_withdraw.iele.json", false)
	debugElrondTest("agar_min_v1/balanceOf.iele.json", false)
	debugElrondTest("agar_min_v1/withdraw_Ok.iele.json", false)
	debugElrondTest("agar_min_v1/withdraw_TooMuch.iele.json", false)
	debugElrondTest("agar_min_v1/rewardWinner.iele.json", false)
	debugElrondTest("agar_min_v1/rewardWinner_Last.iele.json", false)
	debugElrondTest("agar_min_v1/rewardAndSendToWallet.iele.json", false)
	debugElrondTest("agar_min_v1/exceptions.iele.json", false)
}

func debugElrondTest(testFile string, tracePretty bool) {
	err := controller.RunSingleIeleTest(
		filepath.Join(elrondTestRoot, testFile),
		newElrondIeleTestExecutor().SetTracePretty(tracePretty))

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

func debugElrondTestnetTest(testFile string, tracePretty bool) {
	err := controller.RunSingleIeleTest(
		filepath.Join(elrondTestRoot, testFile),
		newElrondTestnetIeleTestExecutor(tracePretty))

	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
}

package main

import (
	"fmt"
	"path"

	eptest "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/endpointtest"
)

func main() {

	//debugTest("tests/iele/danse/factorial/factorial_positive.iele.json")
	//debugTest("tests/iele/albe/factorial/factorial_positive.iele.json")
	//debugTest("tests/iele/albe/ERC20/approve_Caller-Positive.iele.json")
	//debugTest("tests/iele/albe/unit/blockhash.iele.json")
	//debugTest("tests/iele/albe/unit/exceptions.iele.json")
	//debugTest("tests/iele/albe/forwarder/create.iele.json")
	debugTest("tests/iele/albe/ERC20/approve_Caller-Positive.iele.json")

}

func debugTest(testFile string) {
	err := eptest.RunJSONTest(path.Join(ieleTestRoot, testFile))
	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		fmt.Printf("ERROR: %s", err.Error())
	}
}

package main

import (
	"fmt"
	"path"

	endpoint "github.com/ElrondNetwork/elrond-vm/iele-node/endpoint"
)

func main() {

	debugTest("tests/iele/danse/factorial/factorial_positive.iele.json")
	//debugTest("tests/iele/albe/factorial/factorial_positive.iele.json")
	//debugTest("tests/iele/danse/forwarder/create.iele.json")
	//debugTest("tests/iele/albe/ERC20/approve_Caller-Positive.iele.json")

}

func debugTest(testFile string) {
	err := endpoint.RunJSONTest(path.Join(ieleTestRoot, testFile))
	if err == nil {
		fmt.Println("SUCCESS")
	} else {
		panic(err)
	}
}

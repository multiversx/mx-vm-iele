// File provided by the K Framework Go backend. Timestamp: 2019-05-20 22:40:30.522

package main

import (
	interpreter "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestinginterpreter"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		panic("Argument expected. First argument should be the program to execute.")
	}
	execFileName := os.Args[1]
    options := interpreter.ExecuteOptions{TracePretty: false, Verbose: true}
	for _, flag := range os.Args[2:] {
		if flag == "--trace" {
			options.TracePretty = true
		}
	}

	interpreter.ExecuteSimple("../", execFileName, options)
}

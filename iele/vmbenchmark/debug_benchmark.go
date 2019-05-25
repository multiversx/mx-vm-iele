package main

import (
	"go/build"
	"path"

	eiele "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/endpoint"

	einterpreter "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestinginterpreter"
)

var elrondTestRoot = path.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele/elrond/tests/")

func main() {
	eiele.InterpreterOptions = &einterpreter.ExecuteOptions{
		TracePretty: false,
		TraceKPrint: false,
		Verbose:     false,
		MaxSteps:    0,
	}
	benchmarkManyErc20SimpleTransfers(nil, 1)
}

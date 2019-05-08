package ieletestinginterpreter

import (
    "fmt"
	m "github.com/ElrondNetwork/elrond-vm/iele-node/iele-testing-kompiled/ieletestingmodel"
)

var verbose bool = true

func warn(message string) {
    if verbose {
        fmt.Printf("Warning: %s\n", message)
    }
}

// helps us deal with unused variables in some situations
func doNothing(c m.K) {
}

// can be handy when debugging
func debugPrint(c m.K) {
	fmt.Println(m.PrettyPrint(c))
}

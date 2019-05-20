// File provided by the K Framework Go backend. Timestamp: 2019-05-20 22:38:10.632

package ieletestinginterpreter

import (
    "fmt"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
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

// DebugPrint ... prints a K item to console, useful for debugging
func DebugPrint(c m.K) {
	fmt.Println(m.PrettyPrint(c))
}

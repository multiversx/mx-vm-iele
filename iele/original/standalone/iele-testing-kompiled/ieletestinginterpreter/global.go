// File provided by the K Framework Go backend. Timestamp: 2019-06-07 19:55:22.205

package ieletestinginterpreter

import (
    "fmt"
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
)

func (i *Interpreter) warn(message string) {
    if i.Verbose {
        fmt.Printf("Warning: %s\n", message)
    }
}

// helps us deal with unused variables in some situations
func doNothing(c m.K) {
}

// DebugPrint ... prints a K item to console, useful for debugging
func (i *Interpreter) DebugPrint(c m.K) {
	fmt.Println(i.Model.PrettyPrint(c))
}

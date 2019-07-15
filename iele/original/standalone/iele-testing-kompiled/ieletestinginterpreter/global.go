// File provided by the K Framework Go backend. Timestamp: 2019-07-15 11:26:24.140

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
func doNothing(c m.KReference) {
}

// DebugPrint ... prints a K item to console, useful for debugging
func (i *Interpreter) DebugPrint(info string, c m.KReference) {
	fmt.Printf("debug %s: %s\n", info, i.Model.PrettyPrint(c))
}

// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:10:37.856

package ieletestinginterpreter

import (
    "fmt"
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
)

func (i *Interpreter) warn(message string) {
    if i.Verbose {
        fmt.Printf("Warning: %s\n", message)
    }
}

// helps us deal with unused variables in some situations
func doNothing(c m.KReference) {
}

// tricks the compiler to stop complaining about unused vars/boolVars variable in some cases
// never gets called
func doNothingWithVars(varsLen, boolVarLen int) {
}

// DebugPrint ... prints a K item to console, useful for debugging
func (i *Interpreter) DebugPrint(info string, c m.KReference) {
	fmt.Printf("debug (step %d) %s: %s\n", i.currentStep, info, i.Model.PrettyPrint(c))
}

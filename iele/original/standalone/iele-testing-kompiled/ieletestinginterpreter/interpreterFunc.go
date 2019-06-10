// File provided by the K Framework Go backend. Timestamp: 2019-06-07 19:55:22.205

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
)

// GetNrSteps yields how many steps were executed until now from the start of the last execution
func (i *Interpreter) GetNrSteps() int {
    return i.currentStep
}

// GetLastState yields the current (last) state of the interpreter
func (i *Interpreter) GetState() m.K {
     return i.state
}

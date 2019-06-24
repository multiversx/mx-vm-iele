// File generated by the K Framework Go backend. Timestamp: 2019-06-24 20:19:21.575

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
	krypto "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/hookadapter/krypto"
)

// Interpreter is a container with a reference to model and basic options
type Interpreter struct {
	Model         *m.ModelState
	currentStep   int
	MaxSteps      int
	state         m.K
	traceHandlers []traceHandler
	Verbose       bool

	kryptoRef *krypto.Krypto
}
// NewInterpreter creates a new interpreter instance
func NewInterpreter(kryptoRef *krypto.Krypto) *Interpreter {
	model := &m.ModelState{}
	model.Init()

	return &Interpreter {
		Model:         model,
		MaxSteps:      0,
		currentStep:   -1, // meaning that no processing started yet
		state:         nil,
		traceHandlers: nil,
		Verbose:       false,
		kryptoRef: kryptoRef,
	}
}

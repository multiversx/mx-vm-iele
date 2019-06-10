// File provided by the K Framework Go backend. Timestamp: 2019-06-07 19:43:22.780

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

var inputBackup []m.K

// saves a copy of the arguments, so we can later check if they changed during the call
func (interpreter *Interpreter) backupInput(args ...m.K) {
	inputBackup = make([]m.K, len(args))
	for i := 0; i < len(args); i++ {
		inputBackup[i] = interpreter.Model.DeepCopy(args[i])
	}
}

// checks that arguments didn't change in the hook
func (interpreter *Interpreter) checkImmutable(t *testing.T, args ...m.K) {
	if len(args) != len(inputBackup) {
		t.Error("Test not set up properly. Should be the same number of parameters as the last backupInput call.")
	}
	for i := 0; i < len(args); i++ {
		if !interpreter.Model.Equals(args[i], inputBackup[i]) {
			t.Errorf("Input state changed! Got:%s Want:%s", interpreter.Model.PrettyPrint(args[i]), interpreter.Model.PrettyPrint(inputBackup[i]))

		}
	}
}

// newTestInterpreter provides an interpreter for the unit tests
// does not initialize external hooks, even if they exist in the project
// do not make public, the only public constructor should be the one in interpreterDef.go
func newTestInterpreter() *Interpreter {
	model := &m.ModelState{}
	model.Init()

	return &Interpreter {
		Model:         model,
		MaxSteps:      0,
		currentStep:   -1,
		state:         nil,
		traceHandlers: nil,
		Verbose:       false,
	}
}

// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:25:08.138

package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele/elrond/standalone/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

var inputBackup []m.KReference

// saves a copy of the arguments, so we can later check if they changed during the call
func (interpreter *Interpreter) backupInput(args ...m.KReference) {
	inputBackup = make([]m.KReference, len(args))
	for i := 0; i < len(args); i++ {
		inputBackup[i] = interpreter.Model.DeepCopy(args[i])
	}
}

// checks that arguments didn't change in the hook
func (interpreter *Interpreter) checkImmutable(t *testing.T, args ...m.KReference) {
	if len(args) != len(inputBackup) {
		t.Error("Test not set up properly. Should be the same number of parameters as the last backupInput call.")
	}
	for i := 0; i < len(args); i++ {
		copyAgain := interpreter.Model.DeepCopy(args[i])
		if !interpreter.Model.Equals(copyAgain, inputBackup[i]) {
			t.Errorf("Input state changed! Got:%s Want:%s", interpreter.Model.PrettyPrint(args[i]), interpreter.Model.PrettyPrint(inputBackup[i]))

		}
	}
}

// newTestInterpreter provides an interpreter for the unit tests
// does not initialize external hooks, even if they exist in the project
// do not make public, the only public constructor should be the one in interpreterDef.go
func newTestInterpreter() *Interpreter {
	return &Interpreter{
		Model:         m.NewModel(),
		MaxSteps:      0,
		currentStep:   -1,
		state:         m.NullReference,
		traceHandlers: nil,
		Verbose:       false,
	}
}

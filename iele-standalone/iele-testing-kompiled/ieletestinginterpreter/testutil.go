package ieletestinginterpreter

import (
	m "github.com/ElrondNetwork/elrond-vm/iele-standalone/iele-testing-kompiled/ieletestingmodel"
	"testing"
)

var inputBackup []m.K

// saves a copy of the arguments, so we can later check if they changed during the call
func backupInput(args ...m.K) {
	inputBackup = make([]m.K, len(args))
	for i := 0; i < len(args); i++ {
		inputBackup[i] = args[i].DeepCopy()
	}
}

// checks that arguments didn't change in the hook
func checkImmutable(t *testing.T, args ...m.K) {
	if len(args) != len(inputBackup) {
		t.Error("Test not set up properly. Should be the same number of parameters as the last backupInput call.")
	}
	for i := 0; i < len(args); i++ {
		if !args[i].Equals(inputBackup[i]) {
			t.Errorf("Input state changed! Got:%s Want:%s", m.PrettyPrint(args[i]), m.PrettyPrint(inputBackup[i]))

		}
	}
}

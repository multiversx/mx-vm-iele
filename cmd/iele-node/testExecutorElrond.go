package main

import (
	vmi "github.com/ElrondNetwork/elrond-vm-common"
	worldhook "github.com/ElrondNetwork/elrond-vm-util/mock-hook-blockchain"
	cryptohook "github.com/ElrondNetwork/elrond-vm-util/mock-hook-crypto"

	ij "github.com/ElrondNetwork/elrond-vm-util/test-util/ielejson"
	eiele "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/endpoint"
)

type elrondIeleTestExecutor struct {
	tracePretty bool
	logIO       bool
	scheduleVMs map[string]vmi.VMExecutionHandler
	world       *worldhook.BlockchainHookMock
}

func newElrondIeleTestExecutor() *elrondIeleTestExecutor {
	return &elrondIeleTestExecutor{
		tracePretty: false,
		logIO:       false,
		scheduleVMs: make(map[string]vmi.VMExecutionHandler),
		world:       worldhook.NewMock(),
	}
}

// SetTracePretty turns on pretty trace creation, use for debugging only
func (te *elrondIeleTestExecutor) SetTracePretty(tracePretty bool) *elrondIeleTestExecutor {
	te.tracePretty = tracePretty
	return te
}

// SetLogIO causes the VM to print to console all inputs, data from hooks and output
func (te *elrondIeleTestExecutor) SetLogIO() *elrondIeleTestExecutor {
	te.logIO = true
	return te
}

// ProcessCode takes the iele file path, assembles it and yields the bytecode.
func (te *elrondIeleTestExecutor) ProcessCode(testPath string, value string) (string, error) {
	return assembleIeleCode(testPath, value)
}

// Run executes an individual Iele test.
func (te *elrondIeleTestExecutor) Run(test *ij.Test) error {
	scheduleName := test.Network
	schedule, schErr := eiele.ParseSchedule(scheduleName)
	if schErr != nil {
		return schErr
	}
	vm, found := te.scheduleVMs[scheduleName]
	if !found {
		elrondIeleVM := eiele.NewElrondIeleVM(te.world, cryptohook.KryptoHookMockInstance, schedule)
		if te.tracePretty {
			elrondIeleVM.SetTracePretty()
		}
		if te.logIO {
			elrondIeleVM.SetLogIO()
		}
		vm = elrondIeleVM
		te.scheduleVMs[scheduleName] = vm
	}

	err := runTest(test, vm, te.world)
	return err
}

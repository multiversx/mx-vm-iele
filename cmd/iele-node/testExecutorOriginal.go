package main

import (
	vmi "github.com/ElrondNetwork/elrond-vm-common"
	worldhook "github.com/ElrondNetwork/elrond-vm-util/mock-hook-blockchain"
	cryptohook "github.com/ElrondNetwork/elrond-vm-util/mock-hook-crypto"

	ij "github.com/ElrondNetwork/elrond-vm-util/test-util/ielejson"
	oiele "github.com/ElrondNetwork/elrond-vm/iele/original/node/endpoint"
)

type originalIeleTestExecutor struct {
	tracePretty bool
	scheduleVMs map[string]vmi.VMExecutionHandler
	world       *worldhook.BlockchainHookMock
}

func newOriginalIeleTestExecutor(tracePretty bool) *originalIeleTestExecutor {
	return &originalIeleTestExecutor{
		tracePretty: tracePretty,
		scheduleVMs: make(map[string]vmi.VMExecutionHandler),
		world:       worldhook.NewMock(),
	}
}

// ProcessCode takes the iele file path, assembles it and yields the bytecode.
func (te *originalIeleTestExecutor) ProcessCode(testPath string, value string) (string, error) {
	return assembleIeleCode(testPath, value)
}

// Run executes an individual Iele test.
func (te *originalIeleTestExecutor) Run(test *ij.Test) error {
	scheduleName := test.Network
	schedule, schErr := oiele.ParseSchedule(scheduleName)
	if schErr != nil {
		return schErr
	}
	vm, found := te.scheduleVMs[scheduleName]
	if !found {
		originalIeleVM := oiele.NewOriginalIeleVM(te.world, cryptohook.KryptoHookMockInstance, schedule)
		if te.tracePretty {
			originalIeleVM.SetTracePretty()
		}
		vm = originalIeleVM
		te.scheduleVMs[scheduleName] = vm
	}

	err := runTest(test, vm, te.world)
	return err
}

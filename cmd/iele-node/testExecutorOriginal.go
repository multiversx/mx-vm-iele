package main

import (
	vmi "github.com/ElrondNetwork/elrond-vm-common"
	worldhook "github.com/ElrondNetwork/elrond-vm/mock-hook-blockchain"
	cryptohook "github.com/ElrondNetwork/elrond-vm/mock-hook-crypto"

	oiele "github.com/ElrondNetwork/elrond-vm/iele/original/node/endpoint"
	ij "github.com/ElrondNetwork/elrond-vm/iele/test-util/ielejson"
)

type originalIeleTestExecutor struct {
	tracePretty bool
	scheduleVMs map[string]vmi.VMExecutionHandler
	world       *worldhook.BlockchainHookMock
}

func newOriginalIeleTestExecutor(tracePretty bool) *elrondIeleTestExecutor {
	return &elrondIeleTestExecutor{
		tracePretty: tracePretty,
		scheduleVMs: make(map[string]vmi.VMExecutionHandler),
		world:       worldhook.NewMock(),
	}
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
		vm = oiele.NewOriginalIeleVM(te.world, cryptohook.KryptoHookMockInstance, schedule)
		te.scheduleVMs[scheduleName] = vm
	}

	err := runTest(test, vm, te.world)
	return err
}

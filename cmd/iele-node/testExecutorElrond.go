package main

import (
	vmi "github.com/ElrondNetwork/elrond-vm-common"
	worldhook "github.com/ElrondNetwork/elrond-vm/mock-hook-blockchain"
	cryptohook "github.com/ElrondNetwork/elrond-vm/mock-hook-crypto"

	eiele "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/endpoint"
	ij "github.com/ElrondNetwork/elrond-vm/iele/test-util/ielejson"
)

type elrondIeleTestExecutor struct {
	tracePretty bool
	scheduleVMs map[string]vmi.VMExecutionHandler
	world       *worldhook.BlockchainHookMock
}

func newElrondIeleTestExecutor(tracePretty bool) *elrondIeleTestExecutor {
	return &elrondIeleTestExecutor{
		tracePretty: tracePretty,
		scheduleVMs: make(map[string]vmi.VMExecutionHandler),
		world:       worldhook.NewMock(),
	}
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
		vm = elrondIeleVM
		te.scheduleVMs[scheduleName] = vm
	}

	err := runTest(test, vm, te.world)
	return err
}

package main

import (
	worldhook "github.com/ElrondNetwork/elrond-vm/mock-hook-blockchain"
	cryptohook "github.com/ElrondNetwork/elrond-vm/mock-hook-crypto"

	eiele "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/endpoint"
	ij "github.com/ElrondNetwork/elrond-vm/iele/test-util/ielejson"
)

type elrondTestnetIeleTestExecutor struct {
	tracePretty bool
	vm          *eiele.ElrondIeleVM
	world       *worldhook.BlockchainHookMock
}

func newElrondTestnetIeleTestExecutor(tracePretty bool) *elrondTestnetIeleTestExecutor {
	return &elrondTestnetIeleTestExecutor{
		tracePretty: tracePretty,
		vm:          nil,
		world:       worldhook.NewMock(),
	}
}

// Run executes an individual Iele test.
func (te *elrondTestnetIeleTestExecutor) Run(test *ij.Test) error {
	if te.vm == nil {
		te.vm = eiele.NewElrondIeleVM(te.world, cryptohook.KryptoHookMockInstance, eiele.ElrondTestnet)
		if te.tracePretty {
			te.vm.SetTracePretty()
		}
	}

	err := runTest(test, te.vm, te.world)
	return err
}

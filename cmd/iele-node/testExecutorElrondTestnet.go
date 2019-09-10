package main

import (
	worldhook "github.com/ElrondNetwork/elrond-vm-util/mock-hook-blockchain"
	cryptohook "github.com/ElrondNetwork/elrond-vm-util/mock-hook-crypto"

	ij "github.com/ElrondNetwork/elrond-vm-util/test-util/ielejson"
	eiele "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/endpoint"
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

// ProcessCode takes the iele file path, assembles it and yields the bytecode.
func (te *elrondTestnetIeleTestExecutor) ProcessCode(testPath string, value string) (string, error) {
	return assembleIeleCode(testPath, value)
}

// Run executes an individual Iele test.
func (te *elrondTestnetIeleTestExecutor) Run(test *ij.Test) error {
	if te.vm == nil {
		te.vm = eiele.NewElrondIeleVM(
			eiele.TestVMType, eiele.ElrondDefault,
			te.world, cryptohook.KryptoHookMockInstance)
		if te.tracePretty {
			te.vm.SetTracePretty()
		}
	}

	err := runTestElrond(test, te.vm, te.world)
	return err
}

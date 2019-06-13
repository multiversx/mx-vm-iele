package main

import (
	vmi "github.com/ElrondNetwork/elrond-vm-common"
	worldhook "github.com/ElrondNetwork/elrond-vm/mock-hook-blockchain"
	cryptohook "github.com/ElrondNetwork/elrond-vm/mock-hook-crypto"

	ielecommon "github.com/ElrondNetwork/elrond-vm/iele/common"
	eiele "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/endpoint"
	oiele "github.com/ElrondNetwork/elrond-vm/iele/original/node/endpoint"
)

var world = worldhook.NewMock()

type elrondIeleProvider struct {
	tracePretty bool
}

func (p *elrondIeleProvider) GetVM(scheduleName string) (vmi.VMExecutionHandler, error) {
	schedule, err := ielecommon.ParseSchedule(scheduleName)
	if err != nil {
		return nil, err
	}
	vm := eiele.NewElrondIeleVM(world, cryptohook.KryptoHookMockInstance, schedule)
	if p.tracePretty {
		vm.SetTracePretty()
	}
	return vm, nil
}

type originalIeleProvider struct {
	tracePretty bool
}

func (p *originalIeleProvider) GetVM(scheduleName string) (vmi.VMExecutionHandler, error) {
	schedule, err := ielecommon.ParseSchedule(scheduleName)
	if err != nil {
		return nil, err
	}
	vm := oiele.NewOriginalIeleVM(world, cryptohook.KryptoHookMockInstance, schedule)
	if p.tracePretty {
		vm.SetTracePretty()
	}
	return vm, nil
}

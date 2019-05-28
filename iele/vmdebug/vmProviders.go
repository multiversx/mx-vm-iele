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
}

func (elrondIeleProvider) GetVM(scheduleName string) (vmi.VMExecutionHandler, error) {
	schedule, err := ielecommon.ParseSchedule(scheduleName)
	if err != nil {
		return nil, err
	}
	return eiele.NewElrondIeleVM(world, cryptohook.KryptoHookMockInstance, schedule), nil
}

type originalIeleProvider struct {
}

func (originalIeleProvider) GetVM(scheduleName string) (vmi.VMExecutionHandler, error) {
	schedule, err := ielecommon.ParseSchedule(scheduleName)
	if err != nil {
		return nil, err
	}
	return oiele.NewOriginalIeleVM(world, cryptohook.KryptoHookMockInstance, schedule), nil
}

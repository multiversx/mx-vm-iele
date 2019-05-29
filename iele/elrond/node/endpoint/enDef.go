package endpoint

import (
	vmi "github.com/ElrondNetwork/elrond-vm-common"
	ielecommon "github.com/ElrondNetwork/elrond-vm/iele/common"
	blockchain "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/hookadapter/blockchain"
	krypto "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/hookadapter/krypto"
	interpreter "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestinginterpreter"
)

// AddressLength ... Account address length expected by the VM
const AddressLength = 32

// ElrondIeleVM ... Iele VM, Elrond version
type ElrondIeleVM struct {
	schedule     ielecommon.Schedule
	kinterpreter *interpreter.Interpreter
}

// NewElrondIeleVM creates new Elrond Iele VM instance
func NewElrondIeleVM(blockchainHook vmi.BlockchainHook, cryptoHook vmi.CryptoHook, schedule ielecommon.Schedule) *ElrondIeleVM {
	blockchainAdapter := &blockchain.Blockchain{Upstream: blockchainHook}
	kryptoAdapter := &krypto.Krypto{Upstream: cryptoHook}
	kinterpreter := interpreter.NewInterpreter(blockchainAdapter, kryptoAdapter)

	return &ElrondIeleVM{
		schedule:     schedule,
		kinterpreter: kinterpreter,
	}
}

// SetTracePretty turns on pretty trace creation, use for debugging only
func (vm *ElrondIeleVM) SetTracePretty() {
	vm.kinterpreter.SetTracePretty()
}

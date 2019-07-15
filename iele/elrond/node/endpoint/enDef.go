package endpoint

import (
	vmi "github.com/ElrondNetwork/elrond-vm-common"
	blockchain "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/hookadapter/blockchain"
	krypto "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/hookadapter/krypto"
	interpreter "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestinginterpreter"
)

// AddressLength is the account address length expected by the VM.
const AddressLength = 32

// ElrondIeleVM defines an object containing the state of the Iele VM.
// This is the Elrond version.
type ElrondIeleVM struct {
	schedule          Schedule
	blockchainAdapter *blockchain.Blockchain
	kryptoAdapter     *krypto.Krypto
	kinterpreter      *interpreter.Interpreter
}

// NewElrondIeleVM creates new Elrond Iele VM instance
func NewElrondIeleVM(blockchainHook vmi.BlockchainHook, cryptoHook vmi.CryptoHook, schedule Schedule) *ElrondIeleVM {
	blockchainAdapter := &blockchain.Blockchain{
		Upstream:      blockchainHook,
		AddressLength: AddressLength,
	}
	kryptoAdapter := &krypto.Krypto{Upstream: cryptoHook}
	kinterpreter := interpreter.NewInterpreter(blockchainAdapter, kryptoAdapter)

	return &ElrondIeleVM{
		schedule:          schedule,
		blockchainAdapter: blockchainAdapter,
		kryptoAdapter:     kryptoAdapter,
		kinterpreter:      kinterpreter,
	}
}

// ClearVMState resets the VM state without freeing up the memory,
// so the same memory can be reused on the next execution.
func (vm *ElrondIeleVM) ClearVMState() {
	vm.kinterpreter.Model.Clear()
}

// SetTracePretty turns on pretty trace creation, use for debugging only
func (vm *ElrondIeleVM) SetTracePretty() {
	vm.kinterpreter.SetTracePretty()
}

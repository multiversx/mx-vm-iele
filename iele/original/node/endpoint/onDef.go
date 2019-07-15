package endpoint

import (
	vmi "github.com/ElrondNetwork/elrond-vm-common"
	common "github.com/ElrondNetwork/elrond-vm/iele/common"
	blockchain "github.com/ElrondNetwork/elrond-vm/iele/original/node/hookadapter/blockchain"
	krypto "github.com/ElrondNetwork/elrond-vm/iele/original/node/hookadapter/krypto"
	interpreter "github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled/ieletestinginterpreter"
)

// AddressLength is the account address length expected by the VM.
const AddressLength = 20

// OriginalIeleVM defines an object containing the state of the Iele VM.
// This is the original version.
type OriginalIeleVM struct {
	schedule          common.Schedule
	blockchainAdapter *blockchain.Blockchain
	kryptoAdapter     *krypto.Krypto
	kinterpreter      *interpreter.Interpreter
}

// NewOriginalIeleVM creates new original Iele VM instance
func NewOriginalIeleVM(blockchainHook vmi.BlockchainHook, cryptoHook vmi.CryptoHook, schedule common.Schedule) *OriginalIeleVM {
	blockchainAdapter := &blockchain.Blockchain{
		Upstream:      blockchainHook,
		AddressLength: AddressLength,
	}
	kryptoAdapter := &krypto.Krypto{Upstream: cryptoHook}
	kinterpreter := interpreter.NewInterpreter(blockchainAdapter, kryptoAdapter)

	return &OriginalIeleVM{
		schedule:          schedule,
		blockchainAdapter: blockchainAdapter,
		kryptoAdapter:     kryptoAdapter,
		kinterpreter:      kinterpreter,
	}
}

// ClearVMState resets the VM state without freeing up the memory,
// so the same memory can be reused on the next execution.
func (vm *OriginalIeleVM) ClearVMState() {
	vm.kinterpreter.Model.Clear()
}

// SetTracePretty turns on pretty trace creation, use for debugging only
func (vm *OriginalIeleVM) SetTracePretty() {
	vm.kinterpreter.SetTracePretty()
}

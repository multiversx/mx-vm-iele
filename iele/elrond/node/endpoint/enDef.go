package endpoint

import (
	vmi "github.com/ElrondNetwork/elrond-vm-common"
	blockchain "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/hookadapter/blockchain"
	krypto "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/hookadapter/krypto"
	interpreter "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestinginterpreter"
)

// AddressLength is the account address length expected by the VM.
const AddressLength = 32

// TestVMType is the VM type argument we use in tests.
var TestVMType = []byte{0, 0}

// ElrondIeleVM defines an object containing the state of the Iele VM.
// This is the Elrond version.
type ElrondIeleVM struct {
	schedule          Schedule
	blockchainAdapter *blockchain.Blockchain
	kryptoAdapter     *krypto.Krypto
	kinterpreter      *interpreter.Interpreter
	logIO             bool
}

// NewElrondIeleVM creates new Elrond Iele VM instance
func NewElrondIeleVM(
	vmType []byte,
	schedule Schedule,
	blockchainHook vmi.BlockchainHook,
	cryptoHook vmi.CryptoHook) *ElrondIeleVM {

	blockchainAdapter := &blockchain.Blockchain{
		Upstream:      blockchainHook,
		VMType:        vmType,
		AddressLength: AddressLength,
		LogToConsole:  false,
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

// SetTracePretty turns on pretty trace creation, use for debugging only
func (vm *ElrondIeleVM) SetTracePretty() {
	vm.kinterpreter.SetTracePretty()
}

// SetLogIO causes the VM to print to console all inputs, data from hooks and output
func (vm *ElrondIeleVM) SetLogIO() {
	vm.logIO = true
	vm.blockchainAdapter.LogToConsole = true
}

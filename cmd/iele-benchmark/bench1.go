package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"path/filepath"
	"testing"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	compiler "github.com/ElrondNetwork/elrond-vm/iele/compiler"
	eiele "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/endpoint"
	worldhook "github.com/ElrondNetwork/elrond-vm/mock-hook-blockchain"
	cryptohook "github.com/ElrondNetwork/elrond-vm/mock-hook-crypto"
)

var lastReturnCode vmi.ReturnCode

func benchmarkManyErc20SimpleTransfers(b *testing.B, nrTransfers int) {

	contractPathFilePath := filepath.Join(elrondTestRoot, "iele-examples/erc20_elrond.iele")

	compiledBytes := compiler.AssembleIeleCode(contractPathFilePath)
	decoded, err := hex.DecodeString(string(compiledBytes))
	if err != nil {
		panic(err)
	}

	world := worldhook.NewMock()

	contractAddrHex := "c0879ac700000000000000000000000000000000000000000000000000000000"
	account1AddrHex := "acc1000000000000000000000000000000000000000000000000000000000000"
	account2AddrHex := "acc2000000000000000000000000000000000000000000000000000000000000"

	contractAddr, _ := hex.DecodeString(contractAddrHex)
	account1Addr, _ := hex.DecodeString(account1AddrHex)
	account2Addr, _ := hex.DecodeString(account2AddrHex)

	constractStorage := make(map[string][]byte)
	constractStorage[storageKey("01"+account1AddrHex)] = big.NewInt(2000000000).Bytes()
	constractStorage[storageKey("00")] = big.NewInt(2000000000).Bytes() // total supply

	world.AcctMap.PutAccount(&worldhook.Account{
		Exists:  true,
		Address: contractAddr,
		Nonce:   big.NewInt(0),
		Balance: big.NewInt(0),
		Storage: constractStorage,
		Code:    decoded,
	})

	world.AcctMap.PutAccount(&worldhook.Account{
		Exists:  true,
		Address: account1Addr,
		Nonce:   big.NewInt(0),
		Balance: hexToBigInt("e8d4a51000"),
		Storage: make(map[string][]byte),
		Code:    []byte{},
	})

	world.AcctMap.PutAccount(&worldhook.Account{
		Exists:  true,
		Address: account2Addr,
		Nonce:   big.NewInt(0),
		Balance: hexToBigInt("e8d4a51000"),
		Storage: make(map[string][]byte),
		Code:    []byte{},
	})

	// create the VM and allocate some memory
	vm := eiele.NewElrondIeleVM(world, cryptohook.KryptoHookMockInstance, eiele.ElrondDefault)

	if b != nil { // nil when debugging
		b.ResetTimer()
	}

	for benchMarkRepeat := 0; benchMarkRepeat < 1; benchMarkRepeat++ {
		blHeader := &vmi.SCCallHeader{
			Beneficiary: big.NewInt(0),
			Number:      big.NewInt(int64(benchMarkRepeat)),
			GasLimit:    hexToBigInt("174876e800"),
			Timestamp:   big.NewInt(0),
		}

		for txi := 0; txi < nrTransfers; txi++ {
			vm.ClearVMState()

			input := &vmi.ContractCallInput{
				RecipientAddr: contractAddr,
				Function:      "transfer",
				VMInput: vmi.VMInput{
					CallerAddr: account1Addr,
					Arguments: []*big.Int{
						hexToBigInt(account2AddrHex),
						big.NewInt(1),
					},
					CallValue:   big.NewInt(0),
					GasPrice:    big.NewInt(1),
					GasProvided: hexToBigInt("100000"),
					Header:      blHeader,
				},
			}

			output, err := vm.RunSmartContractCall(input)
			if err != nil {
				panic(err)
			}

			if output.ReturnCode != vmi.Ok {
				panic(fmt.Sprintf("returned non-zero code: %d", output.ReturnCode))
			}

			lastReturnCode = output.ReturnCode
		}
	}
}

func hexToBigInt(hexRepresentation string) *big.Int {
	result, ok := big.NewInt(0).SetString(hexRepresentation, 16)
	if !ok {
		panic("invalid hex: " + hexRepresentation)
	}
	return result
}

func storageKey(hexRepresentation string) string {
	decoded, err := hex.DecodeString(hexRepresentation)
	if err != nil {
		panic("invalid hex: " + hexRepresentation)
	}
	return string(decoded)
}

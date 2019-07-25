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

func benchmarkStaticCall(b *testing.B, contract string, functionName string, args ...*big.Int) {

	contractPathFilePath := filepath.Join(elrondTestRoot, contract)

	compiledBytes := compiler.AssembleIeleCode(contractPathFilePath)
	decoded, err := hex.DecodeString(string(compiledBytes))
	if err != nil {
		panic(err)
	}

	world := worldhook.NewMock()

	contractAddrHex := "c0879ac700000000000000000000000000000000000000000000000000000000"
	account1AddrHex := "acc1000000000000000000000000000000000000000000000000000000000000"

	contractAddr, _ := hex.DecodeString(contractAddrHex)
	account1Addr, _ := hex.DecodeString(account1AddrHex)

	world.AcctMap.PutAccount(&worldhook.Account{
		Exists:  true,
		Address: contractAddr,
		Nonce:   big.NewInt(0),
		Balance: big.NewInt(0),
		Storage: make(map[string][]byte),
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

	// create the VM and allocate some memory
	vm := eiele.NewElrondIeleVM(world, cryptohook.KryptoHookMockInstance, eiele.ElrondDefault)

	repeats := 1
	if b != nil { // nil when debugging
		b.ResetTimer()
		repeats = b.N
	}

	for benchMarkRepeat := 0; benchMarkRepeat < repeats; benchMarkRepeat++ {
		blHeader := &vmi.SCCallHeader{
			Beneficiary: big.NewInt(0),
			Number:      big.NewInt(int64(benchMarkRepeat)),
			GasLimit:    hexToBigInt("174876e800"),
			Timestamp:   big.NewInt(0),
		}

		input := &vmi.ContractCallInput{
			RecipientAddr: contractAddr,
			Function:      functionName,
			VMInput: vmi.VMInput{
				CallerAddr:  account1Addr,
				Arguments:   args,
				CallValue:   big.NewInt(0),
				GasPrice:    big.NewInt(1),
				GasProvided: hexToBigInt("100000000000000000000000000000000000000000000000000000000000000000000"),
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

		//fmt.Println("Returned: " + string(output.ReturnData[0].Bytes()))

		lastReturnCode = output.ReturnCode
	}
}

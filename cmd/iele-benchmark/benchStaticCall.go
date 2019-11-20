package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"path/filepath"
	"testing"

	twos "github.com/ElrondNetwork/big-int-util/twos-complement"
	vmi "github.com/ElrondNetwork/elrond-vm-common"
	worldhook "github.com/ElrondNetwork/elrond-vm-util/mock-hook-blockchain"
	cryptohook "github.com/ElrondNetwork/elrond-vm-util/mock-hook-crypto"
	compiler "github.com/ElrondNetwork/elrond-vm/iele/compiler"
	eiele "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/endpoint"
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
		Nonce:   0,
		Balance: big.NewInt(0),
		Storage: make(map[string][]byte),
		Code:    decoded,
	})

	world.AcctMap.PutAccount(&worldhook.Account{
		Exists:  true,
		Address: account1Addr,
		Nonce:   0,
		Balance: hexToBigInt("100000000"),
		Storage: make(map[string][]byte),
		Code:    []byte{},
	})

	convertedArgs := make([][]byte, len(args))
	for i, arg := range args {
		convertedArgs[i] = twos.ToBytes(arg)
	}

	// create the VM and allocate some memory
	vm := eiele.NewElrondIeleVM(
		eiele.TestVMType, eiele.ElrondDefault,
		world, cryptohook.KryptoHookMockInstance)

	repeats := 1
	if b != nil { // nil when debugging
		b.ResetTimer()
		repeats = b.N
	}

	for benchMarkRepeat := 0; benchMarkRepeat < repeats; benchMarkRepeat++ {
		input := &vmi.ContractCallInput{
			RecipientAddr: contractAddr,
			Function:      functionName,
			VMInput: vmi.VMInput{
				CallerAddr:  account1Addr,
				Arguments:   convertedArgs,
				CallValue:   big.NewInt(0),
				GasPrice:    1,
				GasProvided: 100000000,
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

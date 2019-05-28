package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"path/filepath"
	"testing"

	world "github.com/ElrondNetwork/elrond-vm/callback-blockchain"
	compiler "github.com/ElrondNetwork/elrond-vm/iele/compiler"
	eiele "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/endpoint"
	vmi "github.com/ElrondNetwork/elrond-vm-common"
)

var lastReturnCode *big.Int

func benchmarkManyErc20SimpleTransfers(b *testing.B, nrTransfers int) {

	contractPathFilePath := filepath.Join(elrondTestRoot, "iele-examples/erc20_elrond.iele")

	compiledBytes := compiler.AssembleIeleCode(contractPathFilePath)
	decoded, err := hex.DecodeString(string(compiledBytes))
	if err != nil {
		panic(err)
	}

	ws := world.MakeInMemoryWorldState()
	world.HookWorldState = ws

	contractAddrHex := "c0879ac700000000000000000000000000000000000000000000000000000000"
	account1AddrHex := "acc1000000000000000000000000000000000000000000000000000000000000"
	account2AddrHex := "acc2000000000000000000000000000000000000000000000000000000000000"

	contractAddr, _ := hex.DecodeString(contractAddrHex)
	account1Addr, _ := hex.DecodeString(account1AddrHex)
	account2Addr, _ := hex.DecodeString(account2AddrHex)

	constractStorage := make(map[string]*big.Int)
	constractStorage["1"+account1AddrHex] = big.NewInt(2000000000)
	constractStorage["0"] = big.NewInt(2000000000) // total supply

	ws.AcctMap.PutAccount(&world.Account{
		Address: contractAddr,
		Nonce:   big.NewInt(0),
		Balance: big.NewInt(0),
		Storage: constractStorage,
		Code:    string(decoded),
	})

	ws.AcctMap.PutAccount(&world.Account{
		Address: account1Addr,
		Nonce:   big.NewInt(0),
		Balance: hexToBigInt("e8d4a51000"),
		Storage: make(map[string]*big.Int),
		Code:    "",
	})

	ws.AcctMap.PutAccount(&world.Account{
		Address: account2Addr,
		Nonce:   big.NewInt(0),
		Balance: hexToBigInt("e8d4a51000"),
		Storage: make(map[string]*big.Int),
		Code:    "",
	})

	if b != nil { // nil when debugging
		b.ResetTimer()
	}

	for benchMarkRepeat := 0; benchMarkRepeat < 1; benchMarkRepeat++ {
		blHeader := &vmi.BlockHeader{
			Beneficiary:   big.NewInt(0),
			Difficulty:    big.NewInt(0),
			Number:        big.NewInt(int64(benchMarkRepeat)),
			GasLimit:      hexToBigInt("174876e800"),
			UnixTimestamp: big.NewInt(0),
		}

		for txi := 0; txi < nrTransfers; txi++ {
			input := &vmi.VMInput{
				IsCreate:      false,
				CallerAddr:    account1Addr,
				RecipientAddr: contractAddr,
				InputData:     "",
				Function:      "transfer",
				Arguments: []*big.Int{
					hexToBigInt(account2AddrHex),
					big.NewInt(1),
				},
				CallValue:   big.NewInt(0),
				GasPrice:    big.NewInt(1),
				GasProvided: hexToBigInt("100000"),
				BlockHeader: blHeader,
				Schedule:    vmi.Danse,
			}

			vm := eiele.ElrondIeleVM
			output, err := vm.RunTransaction(input)
			if err != nil {
				panic(err)
			}

			if output.ReturnCode.Sign() != 0 {
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

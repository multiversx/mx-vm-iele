package endpoint

import (
	"encoding/hex"
	"fmt"
	"sort"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
)

func (vm *ElrondIeleVM) logCreateInput(input *vmi.ContractCreateInput) {
	if !vm.logIO {
		return
	}

	fmt.Printf("\nContract create started:")
	fmt.Printf("\n  Caller (sender): %s", hex.EncodeToString(input.CallerAddr))
	fmt.Printf("\n  Arguments: ")
	for arg := range input.Arguments {
		fmt.Printf("%d  ", arg)
	}
	fmt.Printf("\n  Call value: %d", input.CallValue)
	fmt.Printf("\n  Gas provided: %d", input.GasProvided)
	fmt.Printf("\n  Gas price: %d", input.GasPrice)
}

func (vm *ElrondIeleVM) logCallInput(input *vmi.ContractCallInput) {
	if !vm.logIO {
		return
	}

	fmt.Printf("\nContract call started:")
	fmt.Printf("\n  Caller (sender):      %s", hex.EncodeToString(input.CallerAddr))
	fmt.Printf("\n  Recipient (receiver): %s", hex.EncodeToString(input.RecipientAddr))
	fmt.Printf("\n  Function: %s", input.Function)
	fmt.Printf("\n  Arguments: ")
	for i, arg := range input.Arguments {
		fmt.Printf("0x%x", arg)
		if i < len(input.Arguments)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Printf("\n  Call value: %d", input.CallValue)
	fmt.Printf("\n  Gas provided: %d", input.GasProvided)
	fmt.Printf("\n  Gas price: %d", input.GasPrice)
}

func (vm *ElrondIeleVM) printAccountData(accounts []*vmi.OutputAccount) {
	addresses := make([]string, len(accounts))
	accMap := make(map[string]*vmi.OutputAccount)
	for i, acc := range accounts {
		addresses[i] = string(acc.Address)
		accMap[string(acc.Address)] = acc
	}
	sort.Strings(addresses)
	for _, addr := range addresses {
		acc := accMap[addr]
		fmt.Printf("\n    %s", hex.EncodeToString(acc.Address))
		fmt.Printf("\n           balance: %d (0x%x)", acc.Balance, acc.Balance)
		fmt.Printf("\n           nonce:   %d (0x%x)", acc.Nonce, acc.Nonce)
		if len(acc.StorageUpdates) > 0 {
			fmt.Print("\n           storage:")
			var stKeys []string
			stMap := make(map[string]string)
			for _, st := range acc.StorageUpdates {
				var key string
				if len(st.Offset) == 0 {
					key = "0"
				} else {
					key = hex.EncodeToString(st.Offset)
				}
				stKeys = append(stKeys, key)
				stMap[key] = hex.EncodeToString(st.Data)
			}
			sort.Strings(stKeys)
			for _, stKey := range stKeys {
				fmt.Printf("\n               0x%s -> 0x%s", stKey, stMap[stKey])
			}
		}
	}
}

func (vm *ElrondIeleVM) logOutput(output *vmi.VMOutput) {
	if !vm.logIO {
		return
	}

	fmt.Printf("\nTransaction output:")
	fmt.Printf("\n  Return code: %d (%s)", int(output.ReturnCode), output.ReturnCode.String())
	fmt.Printf("\n  Output accounts:")
	vm.printAccountData(output.OutputAccounts)
	fmt.Println()
}

func (vm *ElrondIeleVM) logInputAccounts() {
	if !vm.logIO || !vm.blockchainAdapter.LogToConsole {
		return
	}
	fmt.Printf("\nInput accounts, via hook:")
	inputAccounts := vm.blockchainAdapter.GetInputAccounts()
	vm.printAccountData(inputAccounts)
	fmt.Println()
}

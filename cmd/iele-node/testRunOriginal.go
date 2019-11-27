package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"

	worldhook "github.com/ElrondNetwork/elrond-vm-util/mock-hook-blockchain"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	ij "github.com/ElrondNetwork/elrond-vm-util/test-util/vmtestjson"
)

func runTestOriginal(test *ij.Test, vm vmi.VMExecutionHandler, world *worldhook.BlockchainHookMock) error {
	// reset world
	world.Clear()
	world.Blockhashes = test.BlockHashes

	for _, acct := range test.Pre {
		world.AcctMap.PutAccount(convertAccount(acct))
	}

	//spew.Dump(world.AcctMap)

	for _, block := range test.Blocks {
		world.CurrentTimestamp = block.BlockHeader.Timestamp

		for txIndex, tx := range block.Transactions {
			//fmt.Printf("%d\n", txIndex)
			beforeErr := world.UpdateWorldStateBefore(tx.From, tx.GasLimit, tx.GasPrice)
			if beforeErr != nil {
				return beforeErr
			}

			arguments := make([][]byte, len(tx.Arguments))
			for i, arg := range tx.Arguments {
				arguments[i] = append(arguments[i], arg.ToBytesAlwaysForceSign()...)
			}
			var output *vmi.VMOutput

			if tx.IsCreate {
				input := &vmi.ContractCreateInput{
					ContractCode: []byte(tx.AssembledCode),
					VMInput: vmi.VMInput{
						CallerAddr:  tx.From,
						Arguments:   arguments,
						CallValue:   tx.Value,
						GasPrice:    tx.GasPrice,
						GasProvided: tx.GasLimit,
					},
				}

				var err error
				output, err = vm.RunSmartContractCreate(input)
				if err != nil {
					return err
				}
			} else {
				input := &vmi.ContractCallInput{
					RecipientAddr: tx.To,
					Function:      tx.Function,
					VMInput: vmi.VMInput{
						CallerAddr:  tx.From,
						Arguments:   arguments,
						CallValue:   tx.Value,
						GasPrice:    tx.GasPrice,
						GasProvided: tx.GasLimit,
					},
				}

				var err error
				output, err = vm.RunSmartContractCall(input)
				if err != nil {
					return err
				}
			}

			updErr := world.UpdateAccounts(output.OutputAccounts, output.DeletedAccounts)
			if updErr != nil {
				return updErr
			}

			blResult := block.Results[txIndex]

			// check return code
			expectedStatus := 0
			if blResult.Status != nil {
				expectedStatus = int(blResult.Status.Int64())
			}
			if expectedStatus != int(output.ReturnCode) {
				return fmt.Errorf("result code mismatch. Tx #%d. Want: %d. Have: %d", txIndex, expectedStatus, int(output.ReturnCode))
			}

			// check result
			if len(output.ReturnData) != len(blResult.Out) {
				return fmt.Errorf("result length mismatch. Tx #%d. Want: %s. Have: %s",
					txIndex, ij.ResultAsString(blResult.Out), ij.ResultAsString(output.ReturnData))
			}
			for i, expected := range blResult.Out {
				if !ij.ResultEqual(expected, output.ReturnData[i]) {
					return fmt.Errorf("result mismatch. Tx #%d. Want: %s. Have: %s",
						txIndex, ij.ResultAsString(blResult.Out), ij.ResultAsString(output.ReturnData))
				}
			}

			// check refund
			if blResult.Refund != nil {
				if blResult.Refund.Cmp(output.GasRefund) != 0 {
					return fmt.Errorf("result gas refund mismatch. Want: 0x%x. Have: 0x%x",
						blResult.Refund, output.GasRefund)
				}
			}

			// check gas
			if test.CheckGas && blResult.CheckGas {
				if blResult.Gas != output.GasRemaining {
					return fmt.Errorf("result gas mismatch. Want: %d (0x%x). Got: %d (0x%x)",
						blResult.Gas, blResult.Gas, output.GasRemaining, output.GasRemaining)
				}
			}
			// burned := big.NewInt(0).Sub(tx.GasLimit, output.GasRemaining)
			// fmt.Printf("all: 0x%x  remaining: 0x%x  consumed: 0x%x   refund: 0x%x\n", tx.GasLimit, output.GasRemaining, burned, output.GasRefund)

			// check empty logs, this seems to be the value
			if blResult.IgnoreLogs {
				// nothing, ignore
			} else if len(blResult.LogHash) > 0 {
				// for the old tests we only check if the logs are empty or not
				if blResult.LogHash == "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347" {
					if len(output.Logs) != 0 {
						return fmt.Errorf("empty logs expected. Found: %v", blResult.LogHash)
					}
				} else {
					if len(output.Logs) == 0 {
						return fmt.Errorf("non-empty logs expected")
					}
					for _, log := range output.Logs {
						if !bytes.Equal(log.Address, tx.To) {
							return fmt.Errorf("log address mismatch. Want: %s. Got: %s",
								hex.EncodeToString(tx.To), hex.EncodeToString(log.Address))
						}
					}
				}
				// blResult.LogHash = ""
				// blResult.Logs = nil
				// for _, outLog := range output.Logs {
				// 	blResult.Logs = append(blResult.Logs, convertLogToTestFormat(outLog))
				// }
			} else {
				// this is the real log check
				if len(blResult.Logs) != len(output.Logs) {
					return fmt.Errorf("wrong number of logs. Want:%d. Got:%d",
						len(blResult.Logs), len(output.Logs))
				}
				for i, outLog := range output.Logs {
					testLog := blResult.Logs[i]
					if !bytes.Equal(outLog.Address, testLog.Address) {
						return fmt.Errorf("bad log address. Want:\n%s\nGot:\n%s",
							ij.LogToString(testLog), ij.LogToString(convertLogToTestFormat(outLog)))
					}
					if len(outLog.Topics) != len(testLog.Topics) {
						return fmt.Errorf("wrong number of log topics. Want:\n%s\nGot:\n%s",
							ij.LogToString(testLog), ij.LogToString(convertLogToTestFormat(outLog)))
					}
					for ti := range outLog.Topics {
						if big.NewInt(0).SetBytes(outLog.Topics[ti]).Cmp(
							big.NewInt(0).SetBytes(testLog.Topics[ti])) != 0 {
							return fmt.Errorf("bad log topic. Want:\n%s\nGot:\n%s",
								ij.LogToString(testLog), ij.LogToString(convertLogToTestFormat(outLog)))
						}
					}
					if big.NewInt(0).SetBytes(outLog.Data).Cmp(big.NewInt(0).SetBytes(testLog.Data)) != 0 {
						return fmt.Errorf("bad log data. Want:\n%s\nGot:\n%s",
							ij.LogToString(testLog), ij.LogToString(convertLogToTestFormat(outLog)))
					}
				}

			}

		}
	}

	for worldAcctAddr := range world.AcctMap {
		postAcctMatch := ij.FindAccount(test.PostState, []byte(worldAcctAddr))
		if postAcctMatch == nil {
			return fmt.Errorf("unexpected account address: %s", hex.EncodeToString([]byte(worldAcctAddr)))
		}
	}

	for _, postAcctFromTest := range test.PostState {
		postAcct := convertAccount(postAcctFromTest)
		matchingAcct, isMatch := world.AcctMap[string(postAcct.Address)]
		if !isMatch {
			return fmt.Errorf("account %s expected but not found after running test",
				hex.EncodeToString(postAcct.Address))
		}

		if !bytes.Equal(matchingAcct.Address, postAcct.Address) {
			return fmt.Errorf("bad account address %s", hex.EncodeToString(matchingAcct.Address))
		}

		if matchingAcct.Nonce != postAcct.Nonce {
			return fmt.Errorf("bad account nonce. Account: %s. Want: %d. Have: %d",
				hex.EncodeToString(matchingAcct.Address), postAcct.Nonce, matchingAcct.Nonce)
		}

		if matchingAcct.Balance.Cmp(postAcct.Balance) != 0 {
			return fmt.Errorf("bad account balance. Account: %s. Want: 0x%x. Have: 0x%x",
				hex.EncodeToString(matchingAcct.Address), postAcct.Balance, matchingAcct.Balance)
		}

		if !bytes.Equal(matchingAcct.Code, postAcct.Code) {
			return fmt.Errorf("bad account code. Account: %s. Want: [%s]. Have: [%s]",
				hex.EncodeToString(matchingAcct.Address), postAcct.Code, matchingAcct.Code)
		}

		// compare storages
		allKeys := make(map[string]bool)
		for k := range postAcct.Storage {
			allKeys[k] = true
		}
		for k := range matchingAcct.Storage {
			allKeys[k] = true
		}
		storageError := ""
		for k := range allKeys {
			want := postAcct.StorageValue(k)
			have := matchingAcct.StorageValue(k)
			if !bytes.Equal(want, have) {
				storageError += fmt.Sprintf(
					"\n  for key %s: Want: 0x%s. Have: 0x%s",
					hex.EncodeToString([]byte(k)), hex.EncodeToString(want), hex.EncodeToString(have))
			}
		}
		if len(storageError) > 0 {
			return fmt.Errorf("wrong account storage for account 0x%s:%s",
				hex.EncodeToString(postAcct.Address), storageError)
		}
	}

	return nil
}

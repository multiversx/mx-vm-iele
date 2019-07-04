package endpointtest

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	compiler "github.com/ElrondNetwork/elrond-vm/iele/compiler"
	worldhook "github.com/ElrondNetwork/elrond-vm/mock-hook-blockchain"

	vmi "github.com/ElrondNetwork/elrond-vm-common"
	ij "github.com/ElrondNetwork/elrond-vm/iele/test-util/ielejson"
)

// RunJSONTest ... only playing around for now
func RunJSONTest(testFilePath string, vmp VMProvider, world *worldhook.BlockchainHookMock) error {
	var err error
	testFilePath, err = filepath.Abs(testFilePath)
	if err != nil {
		return err
	}

	// Open our jsonFile
	var jsonFile *os.File
	jsonFile, err = os.Open(testFilePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	top, parseErr := ij.ParseTopLevel(byteValue)
	if parseErr != nil {
		return parseErr
	}

	for _, test := range top {
		testErr := runTest(testFilePath, test, vmp, world)
		if testErr != nil {
			return testErr
		}
	}

	// toPath := strings.Replace(testFilePath, "iele-v2", "iele-v3", 1)
	// fmt.Println(toPath)
	// saveModifiedTest(toPath, top)

	return nil
}

func runTest(testFilePath string, test *ij.Test, vmp VMProvider, world *worldhook.BlockchainHookMock) error {
	// reset world
	world.Clear()
	world.Blockhashes = test.BlockHashes

	testDirPath := filepath.Dir(testFilePath)

	scheduleName := test.Network

	var assErr error
	for _, acct := range test.Pre {
		acct.Code, assErr = assembleIeleCode(testDirPath, acct.Code)
		if assErr != nil {
			return assErr
		}
		world.AcctMap.PutAccount(convertAccount(acct))
	}
	for _, acct := range test.PostState {
		acct.Code, assErr = assembleIeleCode(testDirPath, acct.Code)
		if assErr != nil {
			return assErr
		}
	}

	//spew.Dump(world.AcctMap)

	for _, block := range test.Blocks {
		for txIndex, tx := range block.Transactions {
			// refresh VM
			vm, initErr := vmp.GetVM(scheduleName)
			if initErr != nil {
				return initErr
			}

			//fmt.Printf("%d\n", txIndex)
			beforeErr := world.UpdateWorldStateBefore(tx.From, tx.GasLimit, tx.GasPrice)
			if beforeErr != nil {
				return beforeErr
			}

			var output *vmi.VMOutput

			if tx.IsCreate {
				assembledCode, assErr := assembleIeleCode(testDirPath, tx.ContractCode)
				if assErr != nil {
					return assErr
				}

				input := &vmi.ContractCreateInput{
					ContractCode: []byte(assembledCode),
					VMInput: vmi.VMInput{
						CallerAddr:  tx.From,
						Arguments:   tx.Arguments,
						CallValue:   tx.Value,
						GasPrice:    tx.GasPrice,
						GasProvided: nil,
						Header:      convertBlockHeader(block.BlockHeader),
					},
				}

				g0, g0Err := vm.G0Create(input)
				if g0Err != nil {
					return g0Err
				}
				input.GasProvided = big.NewInt(0).Sub(tx.GasLimit, g0)

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
						Arguments:   tx.Arguments,
						CallValue:   tx.Value,
						GasPrice:    tx.GasPrice,
						GasProvided: nil,
						Header:      convertBlockHeader(block.BlockHeader),
					},
				}

				g0, g0Err := vm.G0Call(input)
				if g0Err != nil {
					return g0Err
				}
				input.GasProvided = big.NewInt(0).Sub(tx.GasLimit, g0)

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
					txIndex, resultAsString(blResult.Out), resultAsString(output.ReturnData))
			}
			for i, expected := range blResult.Out {
				if expected.Cmp(output.ReturnData[i]) != 0 {
					return fmt.Errorf("result mismatch. Tx #%d. Want: %s. Have: %s",
						txIndex, resultAsString(blResult.Out), resultAsString(output.ReturnData))
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
			if test.CheckGas && blResult.Gas != nil {
				if blResult.Gas.Cmp(output.GasRemaining) != 0 {
					return fmt.Errorf("result gas mismatch. Want: %d (0x%x). Got: %d (0x%x)",
						blResult.Gas, blResult.Gas, output.GasRemaining, output.GasRemaining)
				}
			}
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
						if outLog.Topics[ti].Cmp(testLog.Topics[ti]) != 0 {
							return fmt.Errorf("bad log topic. Want:\n%s\nGot:\n%s",
								ij.LogToString(testLog), ij.LogToString(convertLogToTestFormat(outLog)))
						}
					}
					if !bytes.Equal(outLog.Data, testLog.Data) {
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

		if matchingAcct.Nonce.Cmp(postAcct.Nonce) != 0 {
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

// for nicer error messages
func resultAsString(result []*big.Int) string {
	str := "["
	for i, res := range result {
		str += fmt.Sprintf("0x%x", res)
		if i < len(result)-1 {
			str += ", "
		}
	}
	return str + "]"
}

func convertAccount(testAcct *ij.Account) *worldhook.Account {
	storage := make(map[string][]byte)
	for _, stkvp := range testAcct.Storage {
		key := string(stkvp.Key.Bytes())
		storage[key] = stkvp.Value.Bytes()
	}

	return &worldhook.Account{
		Exists:  true,
		Address: testAcct.Address,
		Nonce:   big.NewInt(0).Set(testAcct.Nonce),
		Balance: big.NewInt(0).Set(testAcct.Balance),
		Storage: storage,
		Code:    []byte(testAcct.Code),
	}
}

func convertLogToTestFormat(outputLog *vmi.LogEntry) *ij.LogEntry {
	testLog := ij.LogEntry{
		Address: outputLog.Address,
		Topics:  outputLog.Topics,
		Data:    outputLog.Data,
	}
	return &testLog
}

func convertBlockHeader(testBlh *ij.BlockHeader) *vmi.SCCallHeader {
	return &vmi.SCCallHeader{
		Beneficiary: testBlh.Beneficiary,
		Number:      testBlh.Number,
		GasLimit:    testBlh.GasLimit,
		Timestamp:   testBlh.UnixTimestamp,
	}
}

var zero = big.NewInt(0)

func zeroIfNil(i *big.Int) *big.Int {
	if i == nil {
		return zero
	}
	return i
}

// make the tests run faster, by not repeating code assembly over and over again
var assembledCodeCache = make(map[string]string)

func assembleIeleCode(testPath string, value string) (string, error) {
	if value == "" {
		return "", nil
	}
	if strings.HasPrefix(value, "0x") {
		code, _ := hex.DecodeString(value[2:])
		return string(code), nil
	}

	contractPathFilePath := filepath.Join(testPath, value)

	cached, foundInCache := assembledCodeCache[contractPathFilePath]
	if foundInCache {
		return cached, nil
	}

	compiledBytes := compiler.AssembleIeleCode(contractPathFilePath)
	decoded, err := hex.DecodeString(string(compiledBytes))
	if err != nil {
		return "", err
	}

	result := string(decoded)
	assembledCodeCache[contractPathFilePath] = result
	return result, nil
}

// tool to modify tests
// use with caution
func saveModifiedTest(toPath string, top []*ij.Test) {
	resultJSON := ij.ToJSONString(top)

	err := os.MkdirAll(filepath.Dir(toPath), os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(toPath, []byte(resultJSON), 0644)
	if err != nil {
		panic(err)
	}
}

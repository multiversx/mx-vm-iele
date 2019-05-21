package endpoint

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path"
	"path/filepath"
	"strings"

	world "github.com/ElrondNetwork/elrond-vm/callback-blockchain"
	compiler "github.com/ElrondNetwork/elrond-vm/iele/compiler"
	endpoint "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/endpoint"
)

// RunJSONTest ... only playing around for now
func RunJSONTest(testFilePath string) error {
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

	top, parseErr := parseTopLevel(byteValue)
	if parseErr != nil {
		return parseErr
	}

	for _, test := range top.tests {
		testErr := runTest(testFilePath, test)
		if testErr != nil {
			return testErr
		}
	}

	return nil
}

func runTest(testFilePath string, test *test) error {
	ws := world.MakeInMemoryWorldState()
	ws.Blockhashes = test.blockHashes
	world.HookWorldState = ws

	testDirPath := filepath.Dir(testFilePath)

	schedule, schErr := endpoint.ParseSchedule(test.network)
	if schErr != nil {
		return schErr
	}

	var assErr error
	for _, acct := range test.pre {
		acct.Code, assErr = assembleIeleCode(testDirPath, acct.Code)
		if assErr != nil {
			return assErr
		}
		ws.AcctMap.PutAccount(acct)
	}
	for _, acct := range test.postState {
		acct.Code, assErr = assembleIeleCode(testDirPath, acct.Code)
		if assErr != nil {
			return assErr
		}
	}

	//spew.Dump(ws.AcctMap)

	for _, block := range test.blocks {
		for txIndex, tx := range block.transactions {
			var data, function, dataForGas string
			if tx.isCreate {
				data, assErr = assembleIeleCode(testDirPath, tx.contractCode)
				if assErr != nil {
					return assErr
				}
				function = ""
				dataForGas = data
			} else {
				data = ""
				function = tx.function
				dataForGas = function
			}

			beforeErr := endpoint.UpdateWorldStateBefore(ws, tx.from.Bytes(), tx.gasLimit, tx.gasPrice)
			if beforeErr != nil {
				return beforeErr
			}

			g0, g0Err := endpoint.G0(schedule, tx.isCreate, dataForGas, tx.arguments)
			if g0Err != nil {
				return g0Err
			}
			gasLimit := big.NewInt(0).Sub(tx.gasLimit, g0)

			input := &endpoint.VMInput{
				IsCreate:      tx.isCreate,
				CallerAddr:    tx.from,
				RecipientAddr: tx.to,
				InputData:     data,
				Function:      function,
				Arguments:     tx.arguments,
				CallValue:     tx.value,
				GasPrice:      tx.gasPrice,
				GasProvided:   gasLimit,
				BlockHeader:   block.blockHeader,
				Schedule:      schedule,
			}

			output, err := endpoint.RunTransaction(input)
			if err != nil {
				return err
			}

			updErr := endpoint.UpdateWorldStateAfter(ws, output)
			if updErr != nil {
				return updErr
			}

			blResult := block.results[txIndex]

			// check return code
			if blResult.status.Cmp(output.ReturnCode) != 0 {
				return fmt.Errorf("result code mismatch. Want: %d. Have: %d", blResult.status, output.ReturnCode)
			}

			// check result
			if len(output.ReturnData) != len(blResult.out) {
				return fmt.Errorf("result length mismatch. Want: %v. Have: %v", blResult.out, output.ReturnData)
			}
			for i, expected := range blResult.out {
				if expected.Cmp(output.ReturnData[i]) != 0 {
					return fmt.Errorf("result mismatch. Want: %v. Have: %v", blResult.out, output.ReturnData)
				}
			}

			// check refund
			if blResult.refund.Cmp(output.GasRefund) != 0 {
				return errors.New("result gas refund mismatch")
			}

			// check gas
			if blResult.gas.Cmp(output.GasRemaining) != 0 {
				//return fmt.Errorf("result gas mismatch. Want: %d. Got: %d", blResult.gas, output.GasRemaining)
			}

			// check empty logs, this seems to be the value
			if blResult.logs == "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347" {
				if len(output.Logs) != 0 {
					return fmt.Errorf("empty logs expected. Found: %v", blResult.logs)
				}
			} else {
				if len(output.Logs) == 0 {
					return fmt.Errorf("non-empty logs expected")
				}
				for _, log := range output.Logs {
					if log.Address != tx.to {
						return fmt.Errorf("log address mismatch. Want: %d. Got: %d", tx.to, log.Address)
					}
				}
				// ... we're not currently testing for the actual values (via the rlp hash)

				//spew.Dump(output.Logs)
			}

		}
	}

	for worldAcctAddr := range ws.AcctMap {
		_, postAcctMatches := test.postState[worldAcctAddr]
		if !postAcctMatches {
			return fmt.Errorf("unexpected account address: %s", hex.EncodeToString([]byte(worldAcctAddr)))
		}
	}

	for _, postAcct := range test.postState {
		matchingAcct, isMatch := ws.AcctMap[string(postAcct.Address)]
		if !isMatch {
			return fmt.Errorf("account %s expected but not found after running test", hex.EncodeToString(postAcct.Address))
		}

		if !bytes.Equal(matchingAcct.Address, postAcct.Address) {
			return fmt.Errorf("bad account address %s", hex.EncodeToString(matchingAcct.Address))
		}

		if matchingAcct.Nonce.Cmp(postAcct.Nonce) != 0 {
			return fmt.Errorf("bad account nonce. Want: %d. Have: %d", postAcct.Nonce, matchingAcct.Nonce)
		}

		if matchingAcct.Balance.Cmp(postAcct.Balance) != 0 {
			return fmt.Errorf("bad account balance. Want: %d. Have: %d", postAcct.Balance, matchingAcct.Balance)
		}

		if matchingAcct.Code != postAcct.Code {
			return fmt.Errorf("bad account code. Want: %s. Have: %s", postAcct.Code, matchingAcct.Code)
		}

		// compare storages
		allKeys := make(map[string]bool)
		for k := range postAcct.Storage {
			allKeys[k] = true
		}
		for k := range matchingAcct.Storage {
			allKeys[k] = true
		}
		for k := range allKeys {
			want := postAcct.StorageValue(k)
			have := matchingAcct.StorageValue(k)
			if have.Cmp(want) != 0 {
				return fmt.Errorf("wrong account storage entry for key %s. Want: 0x%x. Have: 0x%x",
					k, want, have)
			}
		}
	}

	return nil
}

// make the tests run faster, by not repeating code assembly over and over again
var assembledCodeCache = make(map[string]string)

func assembleIeleCode(testPath string, value string) (string, error) {
	if value == "" {
		return "", nil
	}
	if strings.HasPrefix(value, "0x") {
		return value, nil
	}

	contractPathFilePath := path.Join(testPath, value)

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

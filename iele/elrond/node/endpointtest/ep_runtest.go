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
	interpreter "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/iele-testing-kompiled/ieletestinginterpreter"

	ij "github.com/ElrondNetwork/elrond-vm/iele/test-util/ielejson"
)

// RunJSONTest ... only playing around for now
func RunJSONTest(testFilePath string, tracePretty bool) error {
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
		testErr := runTest(testFilePath, test, tracePretty)
		if testErr != nil {
			return testErr
		}
	}

	// toPath := strings.Replace(testFilePath, "iele-v1", "iele-v2", 1)
	// fmt.Println(toPath)
	// saveModifiedTest(toPath, top)

	return nil
}

func runTest(testFilePath string, test *ij.Test, tracePretty bool) error {
	if tracePretty {
		// for debugging only
		endpoint.InterpreterOptions = &interpreter.ExecuteOptions{
			TracePretty: true,
			TraceKPrint: false,
			Verbose:     false,
			MaxSteps:    0,
		}
	}

	ws := world.MakeInMemoryWorldState()
	ws.Blockhashes = test.BlockHashes
	world.HookWorldState = ws

	testDirPath := filepath.Dir(testFilePath)

	schedule, schErr := endpoint.ParseSchedule(test.Network)
	if schErr != nil {
		return schErr
	}

	var assErr error
	for _, acct := range test.Pre {
		acct.Code, assErr = assembleIeleCode(testDirPath, acct.Code)
		if assErr != nil {
			return assErr
		}
		ws.AcctMap.PutAccount(convertAccount(acct))
	}
	for _, acct := range test.PostState {
		acct.Code, assErr = assembleIeleCode(testDirPath, acct.Code)
		if assErr != nil {
			return assErr
		}
	}

	//spew.Dump(ws.AcctMap)

	for _, block := range test.Blocks {
		for txIndex, tx := range block.Transactions {
			var data, function, dataForGas string
			if tx.IsCreate {
				data, assErr = assembleIeleCode(testDirPath, tx.ContractCode)
				if assErr != nil {
					return assErr
				}
				function = ""
				dataForGas = data
			} else {
				data = ""
				function = tx.Function
				dataForGas = function
			}

			beforeErr := endpoint.UpdateWorldStateBefore(ws, tx.From, tx.GasLimit, tx.GasPrice)
			if beforeErr != nil {
				return beforeErr
			}

			g0, g0Err := endpoint.G0(schedule, tx.IsCreate, dataForGas, tx.Arguments)
			if g0Err != nil {
				return g0Err
			}
			gasLimit := big.NewInt(0).Sub(tx.GasLimit, g0)

			input := &endpoint.VMInput{
				IsCreate:      tx.IsCreate,
				CallerAddr:    tx.From,
				RecipientAddr: tx.To,
				InputData:     data,
				Function:      function,
				Arguments:     tx.Arguments,
				CallValue:     tx.Value,
				GasPrice:      tx.GasPrice,
				GasProvided:   gasLimit,
				BlockHeader:   convertBlockHeader(block.BlockHeader),
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

			blResult := block.Results[txIndex]

			// check return code
			expectedStatus := zeroIfNil(blResult.Status)
			if expectedStatus.Cmp(output.ReturnCode) != 0 {
				return fmt.Errorf("result code mismatch. Want: 0x%x. Have: 0x%x", expectedStatus, output.ReturnCode)
			}

			// check result
			if len(output.ReturnData) != len(blResult.Out) {
				return fmt.Errorf("result length mismatch. Want: %s. Have: %s",
					resultAsString(blResult.Out), resultAsString(output.ReturnData))
			}
			for i, expected := range blResult.Out {
				if expected.Cmp(output.ReturnData[i]) != 0 {
					return fmt.Errorf("result mismatch. Want: %s. Have: %s",
						resultAsString(blResult.Out), resultAsString(output.ReturnData))
				}
			}

			// check refund
			expectedRefund := zeroIfNil(blResult.Refund)
			if expectedRefund.Cmp(output.GasRefund) != 0 {
				return errors.New("result gas refund mismatch")
			}

			// check gas
			expectedGas := zeroIfNil(blResult.Gas)
			if expectedGas.Cmp(output.GasRemaining) != 0 {
				return fmt.Errorf("result gas mismatch. Want: %d. Got: %d", expectedGas, output.GasRemaining)
				//blResult.Gas = output.GasRemaining
			}

			// check empty logs, this seems to be the value
			if blResult.Logs == "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347" {
				if len(output.Logs) != 0 {
					return fmt.Errorf("empty logs expected. Found: %v", blResult.Logs)
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
				// ... we're not currently testing for the actual values (via the rlp hash)

				//spew.Dump(output.Logs)
			}

		}
	}

	for worldAcctAddr := range ws.AcctMap {
		postAcctMatch := ij.FindAccount(test.PostState, []byte(worldAcctAddr))
		if postAcctMatch == nil {
			return fmt.Errorf("unexpected account address: %s", hex.EncodeToString([]byte(worldAcctAddr)))
		}
	}

	for _, postAcctFromTest := range test.PostState {
		postAcct := convertAccount(postAcctFromTest)
		matchingAcct, isMatch := ws.AcctMap[string(postAcct.Address)]
		if !isMatch {
			return fmt.Errorf("account %s expected but not found after running test",
				hex.EncodeToString(postAcct.Address))
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
		storageError := ""
		for k := range allKeys {
			want := postAcct.StorageValue(k)
			have := matchingAcct.StorageValue(k)
			if have.Cmp(want) != 0 {
				storageError += fmt.Sprintf(
					"\n  for key %s: Want: 0x%x. Have: 0x%x",
					k, want, have)
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

func convertAccount(testAcct *ij.Account) *world.Account {
	storage := make(map[string]*big.Int)
	for _, stkvp := range testAcct.Storage {
		key := stkvp.Key.Text(16)
		storage[key] = stkvp.Value
	}

	return &world.Account{
		Address: testAcct.Address,
		Nonce:   big.NewInt(0).Set(testAcct.Nonce),
		Balance: big.NewInt(0).Set(testAcct.Balance),
		Storage: storage,
		Code:    testAcct.Code,
	}
}

func convertBlockHeader(testBlh *ij.BlockHeader) *endpoint.BlockHeader {
	return &endpoint.BlockHeader{
		Beneficiary:   testBlh.Beneficiary,
		Difficulty:    testBlh.Difficulty,
		Number:        testBlh.Number,
		GasLimit:      testBlh.GasLimit,
		UnixTimestamp: testBlh.UnixTimestamp,
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

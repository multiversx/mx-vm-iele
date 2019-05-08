package endpoint

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path"
	"path/filepath"
	"strings"

	eh "github.com/ElrondNetwork/elrond-vm/callback-blockchain"
	compiler "github.com/ElrondNetwork/elrond-vm/iele-compiler"
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
	ws := eh.MakeInMemoryWorldState()
	eh.HookWorldState = ws

	testDirPath := filepath.Dir(testFilePath)

	var assErr error
	for _, acct := range test.pre {
		acct.Code, assErr = assembleIeleCode(testDirPath, acct.Code)
		if assErr != nil {
			return assErr
		}
		ws.AcctMap.PutAccount(acct)
	}

	for _, block := range test.blocks {
		for txIndex, tx := range block.transactions {
			isCreate := tx.to.Sign() == 0
			var data, function, dataForGas string
			if isCreate {
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

			schedule, schErr := parseSchedule(test.network)
			if schErr != nil {
				return schErr
			}

			g0, g0Err := g0(schedule, isCreate, dataForGas, tx.arguments)
			if g0Err != nil {
				return g0Err
			}
			gasLimit := big.NewInt(0).Sub(tx.gasLimit, g0)

			input := &VMInput{
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

			output, err := RunTransaction(input)
			if err != nil {
				return err
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
				return fmt.Errorf("result gas mismatch. Want: %d. Got: %d", blResult.gas, output.GasRemaining)
			}

			//spew.Dump(output)
		}
	}

	return nil

}

func assembleIeleCode(testPath string, value string) (string, error) {
	if value == "" {
		return "", nil
	}
	if strings.HasPrefix(value, "0x") {
		return value, nil
	}

	contractPathFilePath := path.Join(testPath, value)
	compiledBytes := compiler.AssembleIeleCode(contractPathFilePath)
	decoded, err := hex.DecodeString(string(compiledBytes))
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}

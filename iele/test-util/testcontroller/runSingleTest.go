package endpointtest

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	compiler "github.com/ElrondNetwork/elrond-vm/iele/compiler"

	ij "github.com/ElrondNetwork/elrond-vm/iele/test-util/ielejson"
)

// RunSingleIeleTest parses and prepares test, then calls testCallback.
func RunSingleIeleTest(testFilePath string, testExecutor IeleTestExecutor) error {
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
		assembleErr := assembleIeleCodesInTest(testFilePath, test)
		if assembleErr != nil {
			return assembleErr
		}
		testErr := testExecutor.Run(test)
		if testErr != nil {
			return testErr
		}
	}

	// toPath := strings.Replace(testFilePath, "iele-v2", "iele-v3", 1)
	// fmt.Println(toPath)
	// saveModifiedTest(toPath, top)

	return nil
}

func assembleIeleCodesInTest(testFilePath string, test *ij.Test) error {
	testDirPath := filepath.Dir(testFilePath)

	var assErr error
	for _, acct := range test.Pre {
		acct.Code, assErr = assembleIeleCode(testDirPath, acct.Code)
		if assErr != nil {
			return assErr
		}
	}
	for _, acct := range test.PostState {
		acct.Code, assErr = assembleIeleCode(testDirPath, acct.Code)
		if assErr != nil {
			return assErr
		}
	}

	for _, block := range test.Blocks {
		for _, tx := range block.Transactions {
			if tx.IsCreate {
				tx.AssembledCode, assErr = assembleIeleCode(testDirPath, tx.ContractCode)
				if assErr != nil {
					return assErr
				}
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

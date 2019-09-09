package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"

	cryptohook "github.com/ElrondNetwork/elrond-vm-util/mock-hook-crypto"
	oj2k "github.com/ElrondNetwork/elrond-vm-util/test-util/orderedjson2kast"
	krypto "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/hookadapter/krypto"
	interpreter "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestinginterpreter"
	m "github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled/ieletestingmodel"
)

type gasMode string

const (
	gasModeVMTests gasMode = "`VMTESTS_IELE-CONSTANTS`(.KList)"
	gasModeNormal  gasMode = "`NORMAL`(.KList)"
)

// where to find the tests to run
var ieleTestRoot = path.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele/original/tests/")

var kryptoAdapter = &krypto.Krypto{Upstream: cryptohook.KryptoHookMockInstance}

var executionError error

// runTest runs one individual *.iele.json test
func runTest(testFilePath string, testGasMode gasMode, tracePretty bool, b *testing.B) error {
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

	kast, err := oj2k.ConvertOrderedJSONToKast(byteValue, testFilePath, assembleIeleCode)
	if err != nil {
		return err
	}

	//$interpreter "$(dirname "$0")/.build/standalone/iele-testing-kompiled/realdef.cma" -c PGM "$kast" textfile -c SCHEDULE '`DEFAULT_IELE-GAS`(.KList)' text -c MODE '`VMTESTS_IELE-CONSTANTS`(.KList)' text --output-file "$output"
	//$interpreter "$(dirname "$0")/.build/standalone/iele-testing-kompiled/realdef.cma" -c PGM "$kast" textfile -c SCHEDULE '`DEFAULT_IELE-GAS`(.KList)' text -c MODE '`NORMAL`(.KList)' text --output-file "$output" "$@"

	kastInitMap := make(map[string][]byte)
	kastInitMap["PGM"] = []byte(kast)
	kastInitMap["SCHEDULE"] = []byte("`DEFAULT_IELE-GAS`(.KList)")
	kastInitMap["MODE"] = []byte(string(testGasMode))

	kinterpreter := interpreter.NewInterpreter(kryptoAdapter)
	if tracePretty {
		kinterpreter.SetTracePretty()
	}

	repeat := 1
	if b != nil {
		// only when running benchmarks
		b.ResetTimer()
		repeat = b.N
	}

	for i := 0; i < repeat; i++ {
		// execution itself
		executionError = kinterpreter.Execute(kastInitMap)
	}

	if executionError != nil {
		return executionError
	}

	finalState := kinterpreter.GetState()
	nrSteps := kinterpreter.GetNrSteps()

	if !isExitCodeZero(finalState, kinterpreter) {
		return fmt.Errorf(
			"test failed, excution returned non-zero exit code.\nNr. steps performed: %d\nFinal state:\n%s\nTest path:%s",
			nrSteps, kinterpreter.Model.PrettyPrint(finalState), testFilePath)
	}

	if !isKCellEmpty(finalState, kinterpreter) {
		return fmt.Errorf(
			"test failed, K cell not empty in the end.\nNr. steps performed: %d\nFinal state:\n%s\nTest path:%s",
			nrSteps, kinterpreter.Model.PrettyPrint(finalState), testFilePath)
	}

	return nil
}

func isExitCodeZero(c m.KReference, kinterpreter *interpreter.Interpreter) bool {
	if generatedTop, t := kinterpreter.Model.GetKApplyObject(c); t && generatedTop.Label == m.ParseKLabel("<generatedTop>") { // `<generatedTop>`(`<k>`(...
		if exitCodeCell, t := kinterpreter.Model.GetKApplyObject(generatedTop.List[2]); t && exitCodeCell.Label == m.ParseKLabel("<exit-code>") && len(exitCodeCell.List) == 1 { // `<exit-code>`(...
			return kinterpreter.Model.Equals(exitCodeCell.List[0], m.IntZero)
		}

	}

	return false
}

func isKCellEmpty(c m.KReference, kinterpreter *interpreter.Interpreter) bool {
	if generatedTop, t := kinterpreter.Model.GetKApplyObject(c); t && generatedTop.Label == m.ParseKLabel("<generatedTop>") { // `<generatedTop>`(`<k>`(...
		if kcell, t := kinterpreter.Model.GetKApplyObject(generatedTop.List[0]); t && kcell.Label == m.ParseKLabel("<k>") && len(kcell.List) == 1 { // `<k>`(...
			return kinterpreter.Model.KSequenceIsEmpty(kcell.List[0])
		}

	}

	return false
}

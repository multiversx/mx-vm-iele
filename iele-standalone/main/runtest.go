package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	interpreter "github.com/ElrondNetwork/elrond-vm/iele-standalone/iele-testing-kompiled/ieletestinginterpreter"
	m "github.com/ElrondNetwork/elrond-vm/iele-standalone/iele-testing-kompiled/ieletestingmodel"
	oj2k "github.com/ElrondNetwork/elrond-vm/iele-standalone/orderedjson2kast"
)

type gasMode string

const (
	gasModeVMTests gasMode = "`VMTESTS_IELE-CONSTANTS`(.KList)"
	gasModeNormal  gasMode = "`NORMAL`(.KList)"
)

// where to find the tests to run
var ieleTestRoot = path.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele-tests/")

// runTest ... runs one individual *.iele.json test
func runTest(testFilePath string, testGasMode gasMode, tracePretty bool, verbose bool) error {
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

	kast, err := oj2k.ConvertOrderedJSONToKast(byteValue, testFilePath)
	if err != nil {
		return err
	}

	//$interpreter "$(dirname "$0")/.build/standalone/iele-testing-kompiled/realdef.cma" -c PGM "$kast" textfile -c SCHEDULE '`DEFAULT_IELE-GAS`(.KList)' text -c MODE '`VMTESTS_IELE-CONSTANTS`(.KList)' text --output-file "$output"
	//$interpreter "$(dirname "$0")/.build/standalone/iele-testing-kompiled/realdef.cma" -c PGM "$kast" textfile -c SCHEDULE '`DEFAULT_IELE-GAS`(.KList)' text -c MODE '`NORMAL`(.KList)' text --output-file "$output" "$@"

	kastInitMap := make(map[string][]byte)
	kastInitMap["PGM"] = []byte(kast)
	kastInitMap["SCHEDULE"] = []byte("`DEFAULT_IELE-GAS`(.KList)")
	kastInitMap["MODE"] = []byte(string(testGasMode))
	options := interpreter.ExecuteOptions{TracePretty: tracePretty, TraceKPrint: false, Verbose: verbose}

	// execution itself
	finalState, nrSteps, execErr := interpreter.Execute(kastInitMap, options)

	if execErr != nil {
		return execErr
	}

	if !isExitCodeZero(finalState) {
		return fmt.Errorf(
			"test failed, excution returned non-zero exit code.\nNr. steps performed: %d\nFinal state:\n%s\nTest path:%s",
			nrSteps, m.PrettyPrint(finalState), testFilePath)
	}

	if !isKCellEmpty(finalState) {
		return fmt.Errorf(
			"test failed, K cell not empty in the end.\nNr. steps performed: %d\nFinal state:\n%s\nTest path:%s",
			nrSteps, m.PrettyPrint(finalState), testFilePath)
	}

	return nil
}

func isExitCodeZero(c m.K) bool {
	if generatedTop, t := c.(*m.KApply); t && generatedTop.Label == m.ParseKLabel("<generatedTop>") { // `<generatedTop>`(`<k>`(...
		if exitCodeCell, t := generatedTop.List[2].(*m.KApply); t && exitCodeCell.Label == m.ParseKLabel("<exit-code>") && len(exitCodeCell.List) == 1 { // `<exit-code>`(...
			return exitCodeCell.List[0].Equals(m.IntZero)
		}

	}

	return false
}

func isKCellEmpty(c m.K) bool {
	if generatedTop, t := c.(*m.KApply); t && generatedTop.Label == m.ParseKLabel("<generatedTop>") { // `<generatedTop>`(`<k>`(...
		if kcell, t := generatedTop.List[0].(*m.KApply); t && kcell.Label == m.ParseKLabel("<k>") && len(kcell.List) == 1 { // `<k>`(...
			if kseq, isKseq := kcell.List[0].(*m.KSequence); isKseq {
				return kseq.IsEmpty()
			}
		}

	}

	return false
}

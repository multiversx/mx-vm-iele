package endpointtest

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	ij "github.com/ElrondNetwork/elrond-vm/iele/test-util/ielejson"
)

// IeleTestExecutor describes a component that can run a Iele VM test.
type IeleTestExecutor interface {
	Run(*ij.Test) error
}

func isExcluded(excludedFilePatterns []string, testPath string, generalTestPath string) bool {
	for _, et := range excludedFilePatterns {
		excludedFullPath := path.Join(generalTestPath, et)
		match, err := filepath.Match(excludedFullPath, testPath)
		if err != nil {
			panic(err)
		}
		if match {
			return true
		}
	}
	return false
}

// RunAllIeleTestsInDirectory walks directory, parses and prepares all .iele.json tests,
// then calls testCallback for each of them.
func RunAllIeleTestsInDirectory(
	generalTestPath string,
	specificTestPath string,
	excludedFilePatterns []string,
	testExecutor IeleTestExecutor) error {

	mainDirPath := path.Join(generalTestPath, specificTestPath)
	var nrPassed, nrFailed, nrSkipped int

	err := filepath.Walk(mainDirPath, func(testFilePath string, info os.FileInfo, err error) error {
		if strings.HasSuffix(testFilePath, ".iele.json") {
			fmt.Printf("Test: %s ... ", shortenTestPath(testFilePath, generalTestPath))
			if isExcluded(excludedFilePatterns, testFilePath, generalTestPath) {
				nrSkipped++
				fmt.Print("  skip\n")
			} else {
				testErr := RunSingleIeleTest(testFilePath, testExecutor)
				if testErr == nil {
					nrPassed++
					fmt.Print("  ok\n")
				} else {
					nrFailed++
					fmt.Print("  FAIL!!!\n")
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	fmt.Printf("Done. Passed: %d. Failed: %d. Skipped: %d.\n", nrPassed, nrFailed, nrSkipped)
	if nrFailed > 0 {
		return errors.New("Some tests failed")
	}

	return nil
}

func shortenTestPath(path string, generalTestPath string) string {
	if strings.HasPrefix(path, generalTestPath+"/") {
		return path[len(generalTestPath)+1:]
	}
	return path
}

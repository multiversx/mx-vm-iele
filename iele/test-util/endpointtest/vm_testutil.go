package endpointtest

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	vmi "github.com/ElrondNetwork/elrond-vm/iele/vm-interface"
)

var excludedTests = []string{
	"tests/VMTests/vmPerformance/*/*",
	"tests/*/*/unit/precompiled.iele.json",
	"tests/*/*/ill-formed/illFormed.iele.json",
	"tests/*/*/ill-formed/illFormed2.iele.json",
	//"tests/*/*/unit/exceptions.iele.json",
}

func isExcluded(testPath string, generalTestPath string) bool {
	for _, et := range excludedTests {
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

// TestAllInDirectory ... walk directory and run all .iele.json tests
func TestAllInDirectory(t *testing.T, generalTestPath string, specificTestPath string, vm vmi.IeleVM) {
	mainDirPath := path.Join(generalTestPath, specificTestPath)
	var nrPassed, nrFailed, nrSkipped int

	err := filepath.Walk(mainDirPath, func(testFilePath string, info os.FileInfo, err error) error {
		if strings.HasSuffix(testFilePath, ".iele.json") {
			fmt.Printf("Test: %s ... ", shortenTestPath(testFilePath, generalTestPath))
			if isExcluded(testFilePath, generalTestPath) {
				nrSkipped++
				fmt.Print("  skip\n")
			} else {
				testErr := RunJSONTest(testFilePath, vm)
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
		t.Error(err)
	}
	fmt.Printf("Done. Passed: %d. Failed: %d. Skipped: %d.\n", nrPassed, nrFailed, nrSkipped)
	if nrFailed > 0 {
		t.Error("Some tests failed")
	}
}

func shortenTestPath(path string, generalTestPath string) string {
	if strings.HasPrefix(path, generalTestPath+"/") {
		return path[len(generalTestPath)+1:]
	}
	return path
}

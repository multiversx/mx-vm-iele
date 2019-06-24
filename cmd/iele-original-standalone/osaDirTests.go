package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

var excludedTests = []string{
	"tests/VMTests/vmPerformance/*/*",
	"tests/*/*/unit/precompiled.iele.json",
}

func isExcluded(testPath string) bool {
	for _, et := range excludedTests {
		excludedFullPath := path.Join(ieleTestRoot, et)
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

func testAllInDirectory(t *testing.T, mainDirPath string, testGasMode gasMode) {
	var nrPassed, nrFailed, nrSkipped int

	err := filepath.Walk(mainDirPath, func(testFilePath string, info os.FileInfo, err error) error {
		if strings.HasSuffix(testFilePath, ".iele.json") {
			fmt.Printf("Test: %s ... ", shortenTestPath(testFilePath))
			if isExcluded(testFilePath) {
				nrSkipped++
				fmt.Print("  skip\n")
			} else {
				var testErr error
				testErr = runTest(testFilePath, testGasMode, false)
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

func shortenTestPath(path string) string {
	if strings.HasPrefix(path, ieleTestRoot+"/") {
		return path[len(ieleTestRoot)+1:]
	}
	return path
}

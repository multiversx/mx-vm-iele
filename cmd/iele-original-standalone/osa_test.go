package main

import (
	"path"
	"testing"
)

func TestVmTests(t *testing.T) {
	dirPath := path.Join(ieleTestRoot, "tests/VMTests")
	testAllInDirectory(t, dirPath, gasModeVMTests)
}

func TestIeleTests(t *testing.T) {
	dirPath := path.Join(ieleTestRoot, "tests/iele")
	testAllInDirectory(t, dirPath, gasModeNormal)
}

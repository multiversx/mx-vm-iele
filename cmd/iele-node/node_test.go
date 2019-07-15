package main

import (
	"testing"

	controller "github.com/ElrondNetwork/elrond-vm/iele/test-util/testcontroller"
)

var excludedTests = []string{
	//"tests/VMTests/vmPerformance/*/*",
	"tests/*/*/unit/precompiled.iele.json",
}

func TestElrondIeleTests(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"tests/iele-v3",
		excludedTests,
		newElrondIeleTestExecutor(false))

	if err != nil {
		t.Error(err)
	}
}

func TestOriginalIeleTests(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		originalTestRoot,
		"tests/iele",
		excludedTests,
		newOriginalIeleTestExecutor(false))

	if err != nil {
		t.Error(err)
	}
}

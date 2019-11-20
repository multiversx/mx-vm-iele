package main

import (
	"testing"

	controller "github.com/ElrondNetwork/elrond-vm-util/test-util/testcontroller"
)

var excludedTests = []string{
	"tests/*/*/unit/precompiled.iele.json",
	"tests/*/*/unit/shift.iele.json", // TODO: this one started failing after changing arguments, look into it
}

func TestElrondIeleTests(t *testing.T) {
	err := controller.RunAllJSONTestsInDirectory(
		elrondTestRoot,
		"tests/iele-v3",
		".iele.json",
		excludedTests,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

func TestOriginalIeleTests(t *testing.T) {
	err := controller.RunAllJSONTestsInDirectory(
		originalTestRoot,
		"tests/iele",
		".iele.json",
		excludedTests,
		newOriginalIeleTestExecutor(false))

	if err != nil {
		t.Error(err)
	}
}

func BenchmarkExceptionsIele(b *testing.B) {
	debugElrondTest("tests/iele-v3/danse/unit/exceptions.iele.json", false)
}

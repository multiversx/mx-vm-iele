package main

import (
	"testing"

	controller "github.com/ElrondNetwork/elrond-vm-util/test-util/testcontroller"
)

func TestAdder(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"adder",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}

}

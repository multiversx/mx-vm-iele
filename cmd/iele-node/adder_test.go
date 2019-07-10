package main

import (
	"testing"

	controller "github.com/ElrondNetwork/elrond-vm/iele/test-util/testcontroller"
)

func TestAdder(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"adder",
		nil,
		newElrondIeleTestExecutor(false))

	if err != nil {
		t.Error(err)
	}

}

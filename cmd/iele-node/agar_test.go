package main

import (
	"testing"

	controller "github.com/ElrondNetwork/elrond-vm/iele/test-util/testcontroller"
)

func TestAgarV1(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"agar_v1",
		nil,
		newElrondIeleTestExecutor(false))

	if err != nil {
		t.Error(err)
	}
}

func TestAgarV2(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"agar_v2",
		nil,
		newElrondIeleTestExecutor(false))

	if err != nil {
		t.Error(err)
	}
}

func TestAgarV2TestnetGas(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"agar_v2",
		nil,
		newElrondTestnetIeleTestExecutor(false))

	if err != nil {
		t.Error(err)
	}
}

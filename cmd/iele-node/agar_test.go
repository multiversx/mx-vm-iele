package main

import (
	"testing"

	controller "github.com/ElrondNetwork/elrond-vm-util/test-util/testcontroller"
)

func TestAgarMinV1(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"agar_min_v1",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

func TestAgarV1(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"agar_v1",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

func TestAgarV2(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"agar_v2",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

func TestAgarV3(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"agar_v3",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

func TestAgarV4(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"agar_v4",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

func TestAgarV4TestnetGas(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"agar_v4",
		nil,
		newElrondTestnetIeleTestExecutor(false))

	if err != nil {
		t.Error(err)
	}
}

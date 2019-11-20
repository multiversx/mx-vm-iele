package main

import (
	"testing"

	controller "github.com/ElrondNetwork/elrond-vm-util/test-util/testcontroller"
)

func TestAgarMinV1(t *testing.T) {
	err := controller.RunAllJSONTestsInDirectory(
		elrondTestRoot,
		"agar_min_v1",
		".iele.json",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

func TestAgarV1(t *testing.T) {
	err := controller.RunAllJSONTestsInDirectory(
		elrondTestRoot,
		"agar_v1",
		".iele.json",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

func TestAgarV2(t *testing.T) {
	err := controller.RunAllJSONTestsInDirectory(
		elrondTestRoot,
		"agar_v2",
		".iele.json",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

func TestAgarV3(t *testing.T) {
	err := controller.RunAllJSONTestsInDirectory(
		elrondTestRoot,
		"agar_v3",
		".iele.json",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

func TestAgarV4(t *testing.T) {
	err := controller.RunAllJSONTestsInDirectory(
		elrondTestRoot,
		"agar_v4",
		".iele.json",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

func TestAgarV4TestnetGas(t *testing.T) {
	err := controller.RunAllJSONTestsInDirectory(
		elrondTestRoot,
		"agar_v4",
		".iele.json",
		nil,
		newElrondTestnetIeleTestExecutor(false))

	if err != nil {
		t.Error(err)
	}
}

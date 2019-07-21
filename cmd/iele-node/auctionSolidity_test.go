package main

import (
	"testing"

	controller "github.com/ElrondNetwork/elrond-vm/iele/test-util/testcontroller"
)

func TestAuctionSolidity(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"agar_v1",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

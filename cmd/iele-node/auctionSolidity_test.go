package main

import (
	"testing"

	controller "github.com/ElrondNetwork/elrond-vm-util/test-util/testcontroller"
)

func TestAuctionSolidity(t *testing.T) {
	err := controller.RunAllIeleTestsInDirectory(
		elrondTestRoot,
		"auction-solidity",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

package main

import (
	"testing"

	controller "github.com/ElrondNetwork/elrond-vm-util/test-util/testcontroller"
)

func TestAuctionSolidity(t *testing.T) {
	err := controller.RunAllJSONTestsInDirectory(
		elrondTestRoot,
		"auction-solidity",
		".iele.json",
		nil,
		newElrondIeleTestExecutor())

	if err != nil {
		t.Error(err)
	}
}

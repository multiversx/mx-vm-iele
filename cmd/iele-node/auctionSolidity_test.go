package main

import (
	"testing"

	eptest "github.com/ElrondNetwork/elrond-vm/iele/test-util/endpointtest"
)

func TestAuctionSolidity(t *testing.T) {
	eptest.TestAllInDirectory(t,
		elrondTestRoot,
		"auction-solidity",
		&elrondIeleProvider{},
		world)
}

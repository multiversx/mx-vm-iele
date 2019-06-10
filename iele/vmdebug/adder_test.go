package main

import (
	"testing"

	eptest "github.com/ElrondNetwork/elrond-vm/iele/test-util/endpointtest"
)

func TestAdder(t *testing.T) {
	eptest.TestAllInDirectory(t,
		elrondTestRoot,
		"adder",
		&elrondIeleProvider{},
		world)
}

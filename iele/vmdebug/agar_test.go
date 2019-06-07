package main

import (
	"testing"

	eptest "github.com/ElrondNetwork/elrond-vm/iele/test-util/endpointtest"
)

func TestAgar(t *testing.T) {
	eptest.TestAllInDirectory(t,
		elrondTestRoot,
		"agar",
		&elrondIeleProvider{},
		world)
}

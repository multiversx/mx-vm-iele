package main

import (
	"testing"

	eptest "github.com/ElrondNetwork/elrond-vm/iele/test-util/endpointtest"
)

func TestAgarV1(t *testing.T) {
	eptest.TestAllInDirectory(t,
		elrondTestRoot,
		"agar_v1",
		&elrondIeleProvider{},
		world)
}

func TestAgarV2(t *testing.T) {
	eptest.TestAllInDirectory(t,
		elrondTestRoot,
		"agar_v2",
		&elrondIeleProvider{},
		world)
}

package main

import (
	"testing"

	eptest "github.com/ElrondNetwork/elrond-vm/iele/test-util/endpointtest"
)

func TestElrondIeleTests(t *testing.T) {
	eptest.TestAllInDirectory(t,
		elrondTestRoot,
		"tests/iele-v2",
		elrondIeleProvider{},
		world)
}

func TestOriginalIeleTests(t *testing.T) {
	eptest.TestAllInDirectory(t,
		originalTestRoot,
		"tests/iele",
		originalIeleProvider{},
		world)
}

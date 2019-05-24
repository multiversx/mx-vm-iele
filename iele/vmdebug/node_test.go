package main

import (
	"testing"

	eiele "github.com/ElrondNetwork/elrond-vm/iele/elrond/node/endpoint"
	oiele "github.com/ElrondNetwork/elrond-vm/iele/original/node/endpoint"
	eptest "github.com/ElrondNetwork/elrond-vm/iele/test-util/endpointtest"
)

func TestElrondIeleTests(t *testing.T) {
	eptest.TestAllInDirectory(t, elrondTestRoot, "tests/iele-v2", eiele.ElrondIeleVM)
}

func TestOriginalIeleTests(t *testing.T) {
	eptest.TestAllInDirectory(t, originalTestRoot, "tests/iele", oiele.OriginalIeleVM)
}

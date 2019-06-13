package main

import (
	"go/build"
	"path/filepath"
)

// where to find the tests to run
var elrondTestRoot = filepath.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele/elrond/tests/")
var originalTestRoot = filepath.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele/original/tests/")

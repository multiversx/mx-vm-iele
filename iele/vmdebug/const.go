package main

import (
	"go/build"
	"path"
)

// where to find the tests to run
var elrondTestRoot = path.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele/elrond/tests/")
var originalTestRoot = path.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele/original/tests/")

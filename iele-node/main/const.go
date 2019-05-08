package main

import (
	"go/build"
	"path"
)

// where to find the tests to run
var ieleTestRoot = path.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele-tests/")

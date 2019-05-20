#!/bin/sh

rm -rf $GOPATH/src/github.com/ElrondNetwork/elrond-vm/iele/original/node/iele-testing-kompiled
go run $GOPATH/src/github.com/ElrondNetwork/elrond-vm/iele/original/node/kompile

rm -rf $GOPATH/src/github.com/ElrondNetwork/elrond-vm/iele/original/standalone/iele-testing-kompiled
go run $GOPATH/src/github.com/ElrondNetwork/elrond-vm/iele/original/standalone/kompile
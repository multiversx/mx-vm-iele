#!/bin/sh

#also run `mvn install -Dcheckstyle.skip -DskipTests` in K Framework

./kompile-clear.sh

go run $GOPATH/src/github.com/ElrondNetwork/elrond-vm/cmd/kompile/iele-original-node
go run $GOPATH/src/github.com/ElrondNetwork/elrond-vm/cmd/kompile/iele-original-standalone

go run $GOPATH/src/github.com/ElrondNetwork/elrond-vm/cmd/kompile/iele-elrond-node
go run $GOPATH/src/github.com/ElrondNetwork/elrond-vm/cmd/kompile/iele-elrond-standalone

go build ./...
go test ./...
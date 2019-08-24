package main

import (
	"go/build"
	"path"
)

var elrondTestRoot = path.Join(build.Default.GOPATH, "src/github.com/ElrondNetwork/elrond-vm/iele/elrond/tests/")

// Useful commands:
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench BenchmarkManyErc20SimpleTransfers1000
// go tool pprof -http=localhost:4444 cpu.prof
// go tool pprof -http=localhost:4444 mem.prof

func main() {
	benchmarkManyErc20SimpleTransfers(nil, 1)
	//benchmarkStaticCall(nil, "iele-examples/factorial.iele", "factorial", big.NewInt(32))
}

package main

import (
	"path"
	"testing"
)

// Useful commands:
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench BenchmarkLoopMul
// go tool pprof -http=localhost:4444 cpu.prof
// go tool pprof -http=localhost:4444 mem.prof

func runVMTest(testFile string) {
	err := runTest(path.Join(ieleTestRoot, testFile), gasModeVMTests, false)
	if err != nil {
		panic(err)
	}
}

func BenchmarkLoopMul(b *testing.B) {
	for n := 0; n < b.N; n++ {
		runVMTest("tests/VMTests/vmPerformance/loop-mul/loop-mul.iele.json")
	}
}

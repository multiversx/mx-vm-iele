package main

import (
	"path"
	"testing"
)

// Useful commands:
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench=.
// go tool pprof -http=localhost:4444 cpu.prof
// go tool pprof -http=localhost:4444 mem.prof

func runVMTest(testFile string, b *testing.B) {
	err := runTest(path.Join(ieleTestRoot, testFile), gasModeVMTests, false, b)
	if err != nil {
		panic(err)
	}
}

// go test -run=^$ -bench=BenchmarkLoopDivAdd10M
func BenchmarkLoopDivAdd10M(b *testing.B) {
	runVMTest("tests/VMTests/vmPerformance/loop-divadd-10M/loop-divadd-10M.iele.json", b)
}

// go test -run=^$ -bench=BenchmarkLoopNop1M
func BenchmarkLoopNop1M(b *testing.B) {
	runVMTest("tests/VMTests/vmPerformance/loop-exp-nop-1M/loop-exp-nop-1M.iele.json", b)
}

func BenchmarkLoopMul(b *testing.B) {
	runVMTest("tests/VMTests/vmPerformance/loop-mul/loop-mul.iele.json", b)
}

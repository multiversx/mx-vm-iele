package main

import (
	"math/big"
	"testing"
)

// go test -bench BenchmarkLoopExpNop1M
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench BenchmarkLoopExpNop1M
func BenchmarkLoopExpNop1M(b *testing.B) {
	benchmarkStaticCall(b,
		"perf-solidity/basic.iele",
		"testNop(int,int,uint)",
		big.NewInt(0),
		big.NewInt(5),
		big.NewInt(1000000))
}

// go test -bench BenchmarkLoopDivAdd10M
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench BenchmarkLoopDivAdd10M
func BenchmarkLoopDivAdd10M(b *testing.B) {
	benchmarkStaticCall(b,
		"perf-solidity/basic.iele",
		"testDivAdd(uint,uint,uint,uint)",
		big.NewInt(100),
		big.NewInt(300),
		big.NewInt(500),
		big.NewInt(100000))
}

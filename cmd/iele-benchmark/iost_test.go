package main

import (
	"math/big"
	"testing"
)

// go test -cpuprofile cpu.prof -memprofile mem.prof -bench BenchmarkCompute5000
func BenchmarkCompute5000(b *testing.B) {
	benchmarkStaticCall(b,
		"iele-examples/iostBenchmarks.iele",
		"calculate",
		big.NewInt(5000))
}

//go test -cpuprofile cpu.prof -memprofile mem.prof -bench BenchmarkFactorialRec32
func BenchmarkFactorialRec32(b *testing.B) {
	benchmarkStaticCall(b,
		"iele-examples/iostBenchmarks.iele",
		"recursiveFactorial",
		big.NewInt(32))
}

//go test -cpuprofile cpu.prof -memprofile mem.prof -bench BenchmarkStrConcat10000
func BenchmarkStrConcat10000(b *testing.B) {
	str := "This is vm benchmark, tell me who is slower"
	strBigInt := big.NewInt(0).SetBytes([]byte(str))
	benchmarkStaticCall(b,
		"iele-examples/iostBenchmarks.iele",
		"strConcat",
		strBigInt,
		big.NewInt(10000))
}

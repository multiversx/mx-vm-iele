package main

import (
	"math/big"
	"testing"
)

// go test -bench Factorial32
func BenchmarkFactorial32(b *testing.B) {
	benchmarkStaticCall(b, "iele-examples/factorial.iele", "factorial", big.NewInt(32))
}

func BenchmarkManyErc20SimpleTransfers1(b *testing.B) {
	benchmarkManyErc20SimpleTransfers(b, 1)
}

func BenchmarkManyErc20SimpleTransfers10(b *testing.B) {
	benchmarkManyErc20SimpleTransfers(b, 10)
}

func BenchmarkManyErc20SimpleTransfers50(b *testing.B) {
	benchmarkManyErc20SimpleTransfers(b, 50)
}

func BenchmarkManyErc20SimpleTransfers100(b *testing.B) {
	benchmarkManyErc20SimpleTransfers(b, 100)
}

func BenchmarkManyErc20SimpleTransfers1000(b *testing.B) {
	benchmarkManyErc20SimpleTransfers(b, 1000)
}

func BenchmarkManyErc20SimpleTransfers15000(b *testing.B) {
	benchmarkManyErc20SimpleTransfers(b, 15000)
}

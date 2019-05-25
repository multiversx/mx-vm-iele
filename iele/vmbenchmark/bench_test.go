package main

import (
	"testing"
)

func BenchmarkManyErc20SimpleTransfers1(b *testing.B) {
	benchmarkManyErc20SimpleTransfers(b, 1)
}

func BenchmarkManyErc20SimpleTransfers100(b *testing.B) {
	benchmarkManyErc20SimpleTransfers(b, 100)
}

func BenchmarkManyErc20SimpleTransfers1000(b *testing.B) {
	benchmarkManyErc20SimpleTransfers(b, 1000)
}

// func BenchmarkManyErc20SimpleTransfers15000(b *testing.B) {
// 	benchmarkManyErc20SimpleTransfers(b, 15000)
// }

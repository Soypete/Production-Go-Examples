package main

import (
	"testing"
)

func Benchmark2Workers(b *testing.B)  { benchmarkWorkers(2, b) }
func Benchmark5Workers(b *testing.B)  { benchmarkWorkers(5, b) }
func Benchmark10Workers(b *testing.B) { benchmarkWorkers(10, b) }
func Benchmark20Workers(b *testing.B) { benchmarkWorkers(20, b) }

func benchmarkWorkers(i int, b *testing.B) {
	// TODO: fill in the benchmarking code
}

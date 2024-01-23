package main

import (
	"sync"
	"testing"
)

func Benchmark1Workers(b *testing.B)  { benchmarkWorkers(1, b) }
func Benchmark2Workers(b *testing.B)  { benchmarkWorkers(2, b) }
func Benchmark3Workers(b *testing.B)  { benchmarkWorkers(3, b) }
func Benchmark4Workers(b *testing.B)  { benchmarkWorkers(4, b) }
func Benchmark5Workers(b *testing.B)  { benchmarkWorkers(5, b) }
func Benchmark10Workers(b *testing.B) { benchmarkWorkers(10, b) }
func Benchmark20Workers(b *testing.B) { benchmarkWorkers(20, b) }
func Benchmark30Workers(b *testing.B) { benchmarkWorkers(30, b) }

func benchmarkWorkers(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		workerPool := &workerPool{
			wg:         new(sync.WaitGroup),
			msgs:       make(chan string),
			numWorkers: i,
			mu:         new(sync.Mutex),
		}
		go workerPool.queueMessages()
		_ = workerPool.run()
	}
}

# Exercise 4 - Benchmarking

Benchmark is a tool often under utilized by people new to go. The [Benchmarking tool]() ships with the go binary and is part of the [go tools]() testing suite. 

[Pprof]() is another tool that ships with the standard library. This tool gives insights to the [memory information]() and profiling of your go app.

In this exercise we will use both tools to analytics our worker pool app.

## Part 1:

Complete the benchmarking test suite in the file `ex-4-bemarking/worker_test.go`. 

Questions: 
1. What information is provided by the benchmark?
1. Do you consider you code efficient?
1. Post the amount of time your code to took execute with 10 workers, the os, and the processor (your can get this information with `go version`)
	- example: 900ns darwin/arm64

## Part 2:

* access pprof
* analyze memory
* are there potential memory leaks?

# Exercise 4 - Testing and Benchmarking

The Go standard library ships with a bunch of testing tools right out of the box. These incluse the [Main](https://pkg.go.dev/testing#hdr-Main) test tool, the [Benchmarking tool](https://pkg.go.dev/testing#hdr-Benchmarks), the [Fuzzing tool](https://pkg.go.dev/testing#hdr-Fuzzing), and the Unit [Testing tool](https://pkg.go.dev/testing#Testing).

In this exercise we will use both tools to analytics our worker pool app.

## Part 1:

Benchmark is a tool often under utilized by people new to go.

Complete the benchmarking test suite in the file `ex-4-bemarking/worker_test.go`.

run your benchmark tests using the command:

```bash
go test -bench=. -benchmem=true -benchtime=20s
```

Questions:

1.  What information is provided by the benchmark?
1.  Do you consider you code efficient?
1.  Post the amount of time your code to took execute with 10 workers, the os, and the processor (your can get this information with `go version`)
    * example: 900ns darwin/arm64

## Part 2:

1.  Run the code coverage tool
2.  write tests to achieve 60% code coverage.

# Exercise 4 - Benchmarking

Benchmark is a tool often under utilized by people new to go. The [Benchmarking tool]() ships with the go binary and is part of the [go tools]() testing suite. 

[Pprof]() is another tool that ships with the standard library. This tool gives insights to the [memory information]() and profiling of your go app.

In this exercise we will use both tools to analytics our worker pool app.

## Part 1:

Complete the benchmarking test suite in the file `ex-4-bemarking/worker_test.go`. 

run your benchmark tests using the command:
```bash
go test -bench=. -benchmem=true -benchtime=20s
```

Questions: 
1. What information is provided by the benchmark?
1. Do you consider you code efficient?
1. Post the amount of time your code to took execute with 10 workers, the os, and the processor (your can get this information with `go version`)
	- example: 900ns darwin/arm64


### additional practice: Run pprof on your machine
Start by installing graphviz on your machine

On Mac:
```bash
brew install graphviz
```

also setup a memory and cpu profiles by adding these two command line flags
```go
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
```

add this code to your main function to write your memory profile to a file that can be processed with the pprof tool.

```go
if *memprofile != "" {
	f, err := os.Create(*memprofile)
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
}
```

add this code to your main function to write your cpu profile to a file that can be processed with the pprof tool.

```go
if *cpuprofile != "" {
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
}
```
Add a local server to your `func main()` for pprof to scrape:

```go
go func() {
	log.Println(http.ListenAndServe("localhost:6060:, nil))
}()
```

run you main.go in one window. In another window access pprof

```bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

explore more commands [docs](pkg.go.dev/net/http/pprof).
## Part 2:

1. Run the code coverage tool
2. write tests to achieve 60% code coverage.

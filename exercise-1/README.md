# Exercise 1 - Draw Go Scheduler

## Part 1 - scheduling a single process go app: 
Using the free and opensource tool [Draw IO app](https://app.diagrams.net/) draw a diagram worker-pool app from the last exercise being scheduled by the Go Scheduler

## Part 2 - Add pprof to your Go App:
Using the provided `main.go` file, add pprof and explore the memory insights of a single process Go app.

If you are completing this on your own, here are some helpful videos:
* [pprof setup](https://youtu.be/vSdOAzrVvaU)
* [pprof cli](https://youtu.be/Fuz3fNg30cU)
* [pprof web ui](https://youtu.be/v6skRrlXsjY)

### Step 1:
add pprof server to your text parsing app.
First add the pprof driver to your app.

```go
import _ "net/http/pprof"
```

_*NOTE*: the "\_" means that the import is added globally as a backend system. This is common for servers, db drivers, etc_

### Step 2: 
add a pprof server as it's own goroutine in your main function.

```go
// run pprof
go func() {
	http.ListenAndServe("localhost:6060", nil)
}()
```

_*NOTE*: When you do a default `ListenAndServe()` to spin up your server, your pprof is open to the public internet. To add protections use a `mux.Server()` for a custom server and you basic security precautions._

### Step 3: 
install [graphviz](https://graphviz.org/download/) on your machine to get the visual insights.

*Mac:* 
```bash
brew install graphviz
```

### Step 4: 
run pprof while your worker-pool is executing

```bash
go tool pprof -http=:18080 http://localhost:6060/debug/pprof/profile?seconds=30
```

In the default graph each node is a function that your program is running. Size and color indicate how much cpu and time each function is taking.

To acces the commandline tool tool run:

``` bash
go tool pprof http://localhost:6060/debug/pprof/allocs
```

in the command line tool you can search for functions like this

```bash
(pprof) list worker
```

The functions will provide insights in the following categories:

* allocs: A sampling of all past heap memory allocations
* heap: A sampling of heap memory allocations of live objects. 
* profile: CPU profile.
* goroutine: Stack traces of all current goroutines.
* block: Stack traces that led to blocking on synchronization primitives
* cmdline: The command line invocation of the current program
* mutex: Stack traces of holders of contended mutexes
* threadcreate: Stack traces that led to the creation of new OS threads
* trace: A trace of execution of the current program.


### Step 5:

Take some time to expore pprof. Be able to answer the following questions:
1. What function takes the most time?
1. What function take the most cpu?
1. What function takes the most memory?
1. Are any funcitons inlined?


## Resources: 
- [Scheduler saga](https://www.youtube.com/watch?v=YHRO5WQGh0k)
- [pprof for beginners](https://captainnobody1.medium.com/a-beginners-guide-to-pprof-optimizing-your-go-code-c0310e59c485)
- [pprof talk](https://www.youtube.com/watch?v=HjzJ5r2D8ZM)
- [pprof docs](https://github.com/google/pprof/tree/main/doc)
- [GC traces](https://www.ardanlabs.com/blog/2019/05/garbage-collection-in-go-part2-gctraces.html)
- [how to pprof](https://dev.to/agamm/how-to-profile-go-with-pprof-in-30-seconds-592a)
- [Two Go Programs, Three Different Profiling Techniques](https://www.youtube.com/watch?v=nok0aYiGiYA)

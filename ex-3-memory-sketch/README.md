# Exercise 3 - Draw Go Scheduler

## Part 1: 
Using the free and opensource tool [Draw IO app](https://app.diagrams.net/) draw a diagram worker-pool app from the last exercise being scheduled by the Go Scheduler

## Part 2:
Run pprof and see what insights are available to you.

### Step 1:
add pprof server to your worker-pool app. First add the pprof driver to your app. 

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

### Step 3: 
install [graphviz](https://graphviz.org/download/) on your machine to get the visual insights.

*Mac:* 
```bash
brew install graphviz
```

### step 4: 
run pprof while your worker-pool is executing

```bash
go tool pprof -http=:18080 http://localhost:6060/debug/pprof/profile?seconds=30
```

In the default graph each node is a function that your program is running. Size and color indicate how much cpu and time each function is taking.

TODO: add commands of things to look for.

to acces the commandline tool tool run
``` bash
go tool pprof http://localhost:6060/debug/pprof/allocs
```

in the command line tool you can search for functions like this

```bash
(pprof) list worker
```

will show you stats on your code.

Definition:
_span_: how long long a function ran for.
_flat_: number of something
_cum_: cumulative


Questions:

- 


## Resources: 
- [Scheduler saga](https://www.youtube.com/watch?v=YHRO5WQGh0k)
- [GC traces](https://www.ardanlabs.com/blog/2019/05/garbage-collection-in-go-part2-gctraces.html)
- [how to pprof](https://dev.to/agamm/how-to-profile-go-with-pprof-in-30-seconds-592a)
- [Two Go Programs, Three Different Profiling Techniques](https://www.youtube.com/watch?v=nok0aYiGiYA)

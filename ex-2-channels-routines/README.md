# Build a worker pool

Using the template provided in `main.go` build a worker pool using what learned about Go Channels and Goroutines. This example will accept number of 
workers and then use the workers to perform a task. It should take between 20-30 min? 
<!-- how long should this take -->

Follow-up questions: 
- is this a parallel or concurrent process?
- do we need channels and go routines to write this program? What are advantages and disadvantages?
- did your linter suggest any code improvements while you were writing this solution? 

## Running code

You can run the following exercise code locally with the following command. Fill in the `func scheduleWorkers()` to queue up the amount
of workers input via command line arguements.

```bash
go run main.go --workers=2 
```

* use go run command
* accept argement for number of workers
* after specified about of time que up workers to print positive greetings. (what should workers do?)

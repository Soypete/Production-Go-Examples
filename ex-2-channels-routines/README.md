# Build a worker pool

Using the template provided in `main.go` build a worker pool using what learned about Go Channels and Goroutines. This example will accept number of 
workers and then use the workers to perform a task. It should take between 10-15 min? 

## Step 1 build the worker pool

Fill in the `func scheduleWorkers()` to queue up the amount of workers input via command line arguements.

## Step 2 running your code

You can run the following exercise code locally with the following command. 

```bash
go run main.go --workers=2 
```
Follow-up questions: 
- is this a parallel or concurrent process?
- do we need channels and go routines to write this program? What are advantages and disadvantages?
- did your linter suggest any code improvements while you were writing this solution? 


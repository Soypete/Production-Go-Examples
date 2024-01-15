# Build a worker pool

Using the template provided in `main.go` build a worker pool using what you learned about go channels and goroutines. This example will accept a number of workers as a flag `--workers` and then use the workers to perform a task. It should take between 10-15 min?

## Step 1 build the worker pool

Fill in the `func (wp *workerPool) run()` to queue up the amount of workers input via command line arguments.

## Step 2 running your code

You can run the following exercise code locally with the following command. 

```bash
go run main.go --workers=2 
```

Follow-up questions: 
1. is this a parallel or concurrent process?
1. do we need channels and go routines to write this program? What are advantages and disadvantages?
1. did you notice any changes in your pprof metrics?
1. do multiple routines give you any advantage?

# Build a worker pool

Build a worker pool using what you learned about go channels and goroutines. This example will accept a number of workers as a flag `--workers` and then use the workers to perform a task. It should take between 10-15 min?

## Step 1 build the worker pool

Refactor your main.go to accept the flag for workers and then spin up as a concurrent process for each worker. Make sure you use the tools in the [sync package](https://pkg.go.dev/sync) to manage the communication of works across your concurrent processes.

If you did not complete the first exercise feel free to use the template [main.go](/main.go) provided in this repo and fill in the missing funtionality.

## Step 2 runing your code

You can run the following exercise code locally with the following command.

```bash
go run main.go --workers=2
```

Follow-up questions:

1.  is this a parallel or concurrent process?
1.  do we need channels and go routines to write this program? What are advantages and disadvantages?
1.  did you notice any changes in your pprof metrics?
1.  do multiple routines give you any advantage?

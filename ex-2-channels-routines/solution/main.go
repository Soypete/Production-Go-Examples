package main

import (
	"flag"
	"fmt"
	"sync"
)

func queueMessage(ch chan string, msg string) {
	ch <- msg
}

// TODO: clean up main
func main() {
	numWorkers := flag.Int("workers", 1, "number of workers")
	flag.Parse()

	wg := new(sync.WaitGroup)
	ch := make(chan string, 10)

	// start the workers in the background and wait for data on the channel
	// we already know the number of workers, we can increase the WaitGroup once
	wg.Add(*numWorkers)
	for i := 0; i < *numWorkers; i++ {
		go worker(i, ch, wg)
	}

	// make this a loop and randomize the words
	queueMessage(ch, "Hello")
	queueMessage(ch, "World")
	queueMessage(ch, "!")
	queueMessage(ch, "I'm")
	queueMessage(ch, "a")
	queueMessage(ch, "worker")
	queueMessage(ch, "!")
	queueMessage(ch, "I'm")
	queueMessage(ch, "a")
	queueMessage(ch, "coder")

	// close the worker channel and signal there won't be any more data
	close(ch)

	// wait for the workers to stop processing and exit
	wg.Wait()
}

func worker(id int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range ch {
		fmt.Printf("Worker %d received %s\n", id, msg)
	}
}

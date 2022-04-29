package main

import (
	"flag"
	"fmt"
	"sync"
)

// add code here run worker pool
func runWorkerPool(ch chan string, wg *sync.WaitGroup, numWorkers int) {

}

// TODO: clean up main
func main() {
	numWorkers := flag.Int("workers", 1, "number of workers")
	flag.Parse()
	wg := new(sync.WaitGroup)

	ch := make(chan string, 10)

	// make this a loop and randomize the words
	queueMessage(ch, "Hello", wg)
	queueMessage(ch, "World", wg)
	queueMessage(ch, "!", wg)
	queueMessage(ch, "I'm", wg)
	queueMessage(ch, "a", wg)
	queueMessage(ch, "worker", wg)
	queueMessage(ch, "!", wg)
	queueMessage(ch, "I'm", wg)
	queueMessage(ch, "a", wg)
	queueMessage(ch, "coder", wg)

	runWorkerPool(ch, wg, *numWorkers)
}

func queueMessage(ch chan string, msg string, wg *sync.WaitGroup) {
	wg.Add(1)
	ch <- msg
}

func worker(id int, ch chan string, wg *sync.WaitGroup) {
	for msg := range ch {
		fmt.Print(msg)
		wg.Done()
	}
}

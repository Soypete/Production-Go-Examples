package main

import (
	"flag"
	"fmt"
	"sync"
)

func queueMessage(ch chan string, msg string, wg *sync.WaitGroup) {
	wg.Add(1)
	ch <- msg
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

	for i := 0; i < *numWorkers; i++ {
		go worker(i, ch, wg)
	}
	wg.Wait()
}

func worker(id int, ch chan string, wg *sync.WaitGroup) {
	for msg := range ch {
		fmt.Printf("Worker %d received %s\n", id, msg)
		wg.Done()
	}
}

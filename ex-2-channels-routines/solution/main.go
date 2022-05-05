package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// add code here run worker pool
func runWorkerPool(ch chan string, wg *sync.WaitGroup, numWorkers int) {

	for i := 0; i < numWorkers; i++ {
		go worker(i, ch, wg)
	}
}

// TODO: clean up main
func main() {
	numWorkers := flag.Int("workers", 1, "number of workers")
	flag.Parse()
	rand.Seed(time.Now().Unix())

	wg := new(sync.WaitGroup)
	ch := make(chan string)

	// start the workers in the background and wait for data on the channel
	// we already know the number of workers, we can increase the WaitGroup once
	wg.Add(*numWorkers)

	runWorkerPool(ch, wg, *numWorkers)
	queueMessages(ch)

	// wait for the workers to stop processing and exit
	wg.Wait()
}

// getMessages gets a slice of messages to process
func getMessages() []string {
	return []string{"Hello", "World", "!", "I'm", "a", "worker"}
}

func queueMessages(ch chan string) {
	msgs := getMessages()
	for _, msg := range msgs {
		ch <- msg
	}

	// close the worker channel and signal there won't be any more data
	close(ch)

}

func worker(id int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range ch {
		fmt.Printf("Worker %d received %s\n", id, msg)

		// simulate work
		length := time.Duration(rand.Int63n(5000))
		time.Sleep(length * time.Millisecond)
	}
}

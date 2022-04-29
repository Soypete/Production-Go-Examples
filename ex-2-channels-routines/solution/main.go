package main

import (
	"flag"
	"fmt"
	"sync"
)

<<<<<<< HEAD
// add code here run worker pool
func runWorkerPool(ch chan string, wg *sync.WaitGroup, numWorkers int) {

	for i := 0; i < numWorkers; i++ {
		go worker(i, ch, wg)
	}

	wg.Wait()
=======
func queueMessage(ch chan string, msg string) {
	ch <- msg
>>>>>>> 7e300842dc5a46b1d21d6c8f4d85a738ec6dc23d
}

// TODO: clean up main
func main() {
	numWorkers := flag.Int("workers", 1, "number of workers")
	flag.Parse()

	wg := new(sync.WaitGroup)
	ch := make(chan string, 10)

<<<<<<< HEAD
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
=======
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
>>>>>>> 7e300842dc5a46b1d21d6c8f4d85a738ec6dc23d
}

func worker(id int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range ch {
<<<<<<< HEAD
		fmt.Print(msg)
		wg.Done()
=======
		fmt.Printf("Worker %d received %s\n", id, msg)
>>>>>>> 7e300842dc5a46b1d21d6c8f4d85a738ec6dc23d
	}
}

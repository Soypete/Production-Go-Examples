package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"sync"
	"time"
)

type workerPool struct {
	numWorkers int
	mu         *sync.Mutex
	msgs       chan string
	wg         *sync.WaitGroup
}

func (wp *workerPool) run() int {
	numWords := 0
	// start the workers in the background and wait for data on the channel
	// we already know the number of workers, we can increase the WaitGroup once
	wp.wg.Add(wp.numWorkers)
	for i := 0; i < wp.numWorkers; i++ {
		go func() {
			defer wp.wg.Done()
			wordsDetected := wp.detectWords()
			wp.mu.Lock()
			defer wp.mu.Unlock()
			numWords = numWords + wordsDetected
		}()
	}

	// wait for the workers to stop processing and exit
	wp.wg.Wait()
	return numWords
}

// getMessages gets a slice of messages to process
func getMessages() []string {
	file, _ := os.ReadFile("datums/melville-moby_dick.txt")
	words := strings.Split(string(file), " ")
	return words
}

// this will block and not close if the len(msgs) is larger than the channel buffer.
func (wp *workerPool) queueMessages() {
	msgs := getMessages()
	for _, msg := range msgs {
		// add messages to string channel
		wp.msgs <- msg
	}

	// close the worker channel and signal there won't be any more data
	close(wp.msgs)
}

func (wp workerPool) detectWords() int {
	var numWordsDetected int
	for word := range wp.msgs {
		// print detected word to slow down processing and run pprof
		fmt.Printf("%s\n", word)
		// simulate work
		length := time.Duration(rand.Int63n(100))
		time.Sleep(length * time.Millisecond)
		// this condition returns words like whale, whaling, whales
		if strings.Contains(word, "whal") {
			wp.mu.Lock()
			numWordsDetected++
			wp.mu.Unlock()
		}
	}
	return numWordsDetected
}

func main() {
	// setup and configs
	numWorkers := flag.Int("workers", 1, "number of workers")
	flag.Parse()
	startTime := time.Now()
	rand.Seed(time.Now().Unix())
	// run pprof
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	workerPool := &workerPool{
		wg:         new(sync.WaitGroup),
		msgs:       make(chan string),
		numWorkers: *numWorkers,
		mu:         new(sync.Mutex),
	}

	// start the workers in the background and wait for data on the channel
	// we already know the number of workers, we can increase the WaitGroup once
	go workerPool.queueMessages()
	numWords := workerPool.run()
	fmt.Printf("Number of words: %d\nTime to process file: %2f seconds", numWords, time.Since(startTime).Seconds())
}

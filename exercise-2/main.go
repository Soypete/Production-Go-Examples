package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
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
	// TODO: add logic for queueing workers based on the number provided in the flag.
	// add the number of words detected by each worker to the totalNumberWorkder
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
		// simulate work
		length := time.Duration(rand.Int63n(50))
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

// export cpu and mem profiles to a file that can be processed by pprof tool
func runPprof(memprofile, cpuprofile string) {

	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if memprofile != "" {
		f, err := os.Create(memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

func main() {
	numWorkers := flag.Int("workers", 1, "number of workers")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile := flag.String("memprofile", "", "write memory profile to `file`")

	flag.Parse()
	startTime := time.Now()

	// run pprof
	runPprof(*cpuprofile, *memprofile)

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

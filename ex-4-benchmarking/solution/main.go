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

func runWorkerPool(ch chan string, wg *sync.WaitGroup, numWorkers int) {

	// start the workers in the background and wait for data on the channel
	// we already know the number of workers, we can increase the WaitGroup once
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			worker(ch, wg)
			wg.Done()
		}()
	}
	// wait for the workers to stop processing and exit
	wg.Wait()
}

// getMessages gets a slice of messages to process
func getMessages() []string {
	// file, _ := os.ReadFile("datums/melville-moby_dick.txt")
	words := strings.Split("We dont want to overload io threads and the runtime while we are benchmarking, so we are addings some others words that are not as much but talk about go and other cool software things", " ")
	return words
}

// this will block and not close if the len(msgs) is larger than the channel buffer.
func queueMessages(ch chan string) {
	msgs := getMessages()
	for _, msg := range msgs {
		ch <- msg
	}

	// close the worker channel and signal there won't be any more data
	close(ch)
}

func worker(ch chan string, wg *sync.WaitGroup) {
	for range ch {
		// print msg
		// fmt.Printf("%s\n", msg)

		// simulate work
		length := time.Duration(rand.Int63n(50))
		time.Sleep(length * time.Millisecond)
	}
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
	rand.Seed(time.Now().Unix())

	// run pprof
	runPprof(*cpuprofile, *memprofile)

	wg := new(sync.WaitGroup)
	ch := make(chan string)

	// start the workers in the background and wait for data on the channel
	// we already know the number of workers, we can increase the WaitGroup once
	go queueMessages(ch)
	runWorkerPool(ch, wg, *numWorkers)
}

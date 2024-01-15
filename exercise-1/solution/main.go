package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"time"
)

// getWords gets a slice of messages to process
func getWords() []string {
	file, _ := os.ReadFile("datums/melville-moby_dick.txt")
	words := strings.Split(string(file), " ")
	return words
}

func detectWords() int {
	msgs := getWords()
	var numWordsDetected int
	for _, word := range msgs {
		// this condition returns words like whale, whaling, whales
		if strings.Contains(word, "whal") {
			numWordsDetected++
			// golang is too powerful, so we have to slow it down to run pprof
			time.Sleep(50 * time.Millisecond)
		}
	}
	return numWordsDetected

}

func main() {
	// run pprof
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	startTime := time.Now()
	// start the workers in the background and wait for data on the channel
	// we already know the number of workers, we can increase the WaitGroup once
	numWords := detectWords()
	fmt.Printf("Number of words: %d\nTime to process file: %2f seconds", numWords, time.Since(startTime).Seconds())
}

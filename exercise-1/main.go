package main

import (
	"fmt"
	"math/rand"
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
		// golang is too powerful, so we have to slow it down to run pprof
		// this 'sleep' is to simlutate work.
		length := time.Duration(rand.Int63n(50))
		time.Sleep(length * time.Millisecond)
		// this condition returns words like whale, whaling, whales
		if strings.Contains(word, "whal") {
			numWordsDetected++
		}
	}
	return numWordsDetected

}

func main() {

	startTime := time.Now()
	// start the workers in the background and wait for data on the channel
	// we already know the number of workers, we can increase the WaitGroup once
	numWords := detectWords()
	fmt.Printf("Number of words: %d\nTime to process file: %2f seconds", numWords, time.Since(startTime).Seconds())
}

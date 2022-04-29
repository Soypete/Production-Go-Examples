package main

import (
	"fmt"
	"sync"
)

// example stolen from blog post: https://yourbasic.org/golang/gotcha-data-race-closure/
func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Print(i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println()

}

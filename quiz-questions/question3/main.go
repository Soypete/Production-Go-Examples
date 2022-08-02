package main

import (
	"fmt"
	"time"
)

// https://go.dev/play/p/NhkljNuluRO credit to @morgahl_ from twitch
func main() {
	for i := range make([]struct{}, 5) {
		go func() {
			fmt.Println(i)
		}()
	}
	// for i := range make([]struct{}, 5) {
	// 	go func(i int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }
	time.Sleep(time.Second)
}

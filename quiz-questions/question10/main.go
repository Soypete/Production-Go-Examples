package main

import "fmt"

// credit to @jnewmano https://go.dev/play/p/aPa-M0u2CSV
func main() {
	str := "Hello, 世界"
	fmt.Println(len(str))
	for i, v := range str {
		fmt.Printf("index:%d rune:%U string:%v Type: %T\n", i, v, string(v), v)
	}
}

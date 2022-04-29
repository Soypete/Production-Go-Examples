package main

import "fmt"

func main() {
	slice1 := []string{"A", "B", "C", "D", "E"}
	slice2 := slice1[2:4]
	slice2[0] = "changed"

	fmt.Printf("Length [%d] Capacity[%d] %v\n", len(slice1), cap(slice1), slice1)
	// fmt.Printf("Length [%d] Capacity[%d] %v", len(slice2), cap(slice2), slice2)
}

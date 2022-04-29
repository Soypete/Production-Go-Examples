package main

import "fmt"

type ExampleStruct struct{}

func (s *ExampleStruct) Do(printThis string) {
	fmt.Printf("Print: %s\n", printThis)
}

type DoerOfThings interface {
	Do(string)
}

func main() {
	s := &ExampleStruct{}

	// this is magical - non readable code ¯\_(ツ)_/¯
	if d, ok := interface{}(s).(DoerOfThings); ok {
		d.Do("we made it")

		// address, Types, value
		fmt.Printf("Address: %p, Type: %T, Value: %v\n", s, s, s)
		fmt.Printf("Address: %p, Type: %T, Value: %v\n", d, d, d)
	}
}

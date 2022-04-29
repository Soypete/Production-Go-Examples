package main

import (
	"fmt"
)

func wrapError(err error) error {
	return fmt.Errorf("wrapping Error: %w", err)
}

var movies0to5 = []string{"Wall-e", "The Lego Movie", "Lilo and Stitch"}
var movies5to10 = []string{"The Lego Movie", "Star Wars", "Home", "Back To the Future"}
var movies10to13 = []string{"Star Wars", "E.T.", "Back To the Future", "Ghostbusters"}
var movies13to17 = []string{"Star Wars", "Guardians of the Galaxy", "Interstellar", "Arrival"}
var movies17up = []string{"Star Wars", "Aliens", "The Matrix", "The Terminator"}

func getMovies(age int) ([]string, error) {
	switch {
	case age >= 17:
		return movies17up, nil
	case age < 17 && age >= 13:
		return movies13to17, nil
	case age < 13 && age >= 10:
		return movies10to13, nil
	case age < 10 && age >= 5:
		return movies5to10, nil
	case age < 0:
		return []string{}, fmt.Errorf("age must be a positive number")
	default:
		return movies0to5, nil
	}
}

func main() {

	go getMovies(-1)
}

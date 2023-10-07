package main

import "fmt"

func main() {
	// age := [10]int {12, 27, 31, 14, 75, 36, 67, 58, 40, 10}
	age := []int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// i already represent the elements of the age array with the use...
	// of range looping through the age array so it is not age[i], but i.
	// _, replaced i or index, but the _, is to avoid unused varable error messages
	// a represents the elements in the age array
	for _, a := range age {
		if a % 2 == 0 {
			fmt.Printf("the age %d is even \n", a)
		} else {
			fmt.Printf("the age %d is odd \n", a)
		}
	}
}

// you use age[i] when you're not looping through the age, ...
// that is if it's only an if statement
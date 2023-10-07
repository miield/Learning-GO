package main

import "fmt"

func main() {

	// arrays are fixed size element type, and cannot be changed
	scores := [4]int { 1, 2, 3, 4}
	fmt.Println(scores)

	// slices are dynamically-sized element type, and are mostly used because they are flexible
	ages := []int{5, 8, 3, 9}
	// modifying slice arrays
	ages[2] = 13
	fmt.Println(ages[2])

	// slice append; adding to slice
	ages = append(ages, 1) // assigns permanently
	fmt.Println(append(ages, 15)) // appends temporaries to slice at execution time
	fmt.Println(ages, len(ages))

	// array range
	names := []string {"Oyindamola", "Efunroye", "Abiola", "Dasola", "Apeke", "Abosede"}
	// starts at length 2, index 1 and stops at length 3 index 3 - 1
	rangeOne := names[1:3] 
	// starts at index 2 and stops at the last index
	rangeTwo := names[2:]
	// starts at index 0 and stops at length 4 index(3) 4 - 1
	rangeThree := names[:4]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	// Ranges can be append as well
}
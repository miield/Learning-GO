package main

import "fmt"

func main() {
	classPosition := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := range classPosition {
		if classPosition[i] % 2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

}
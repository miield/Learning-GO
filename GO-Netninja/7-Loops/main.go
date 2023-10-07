package main 

import (
	"fmt"
)

func main() {
	// GO always use 'for' for all kinds of loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	names := []string {"Oyindamola", "Efunroye", "Abiola", "Dasola", "Apeke", "Abosede"}

	for n := 0; n < len(names); n++ {
		fmt.Println(names[n])
	}
}
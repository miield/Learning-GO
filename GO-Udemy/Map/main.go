package main

import "fmt"

// func main() {
// 	// declare mapping
// 	addr := map[string]string{
// 		"home": "12, Road D, Awesu Igbogbo",
// 		"office": "Web3Bridge, Abadek Igbogbo",
// 	}

// 	fmt.Println(addr)
// }

func main() {

	// create an empty map
	// var addr map[string]string

	// create an empty map
	// colors := make(map[string]string)
	// // add the data to the map
	// colors["white"] = "#000000"
	// colors["black"] = "#001000"
	// colors["blue"] = "#000200"

	// To delete the map items
	// delete(color, "white")

	// create a map: int = key, string = value
	food := map[int]string{
		001: "plantain is sweet",
		002: "egg is delicious",
	}

	// call the printMap function with the colors variable
	printMap(food)
}

// Iterate over the map items
// c = may variable representing the color in the main map
// map[int]string = type key=value pair
func printMap(c map[int]string) {
	for color, hex := range c {
		fmt.Println("Hex value of ", color, " is ", hex)
	}
}

// DONE by me 
// func main() {
// 	colors := map[int]string {
// 		101 : "Yellow",
// 		102 : "Black",
// 		103 : "White",
// 		104 : "Nude",
// 	}

// 	// fmt.Println(colors)

// 	// call function printColors to print the colors
// 	printColors(colors)
// }

// // to iterate over all the colors
// func printColors (c map[int]string) {
// 	for id, color := range c {
// 		fmt.Println(id, "is the color number of ", color)
// 	}
// }
package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangleArea struct {
	height float64
	base float64
}
type squareArea struct {
	sideLength float64
}

func main() {

	// the variables below are used to rep the structs and indirectly
	// connected with the interface
	ta := triangleArea{height: 4, base: 3}
	sa := squareArea{sideLength: 7}

	printArea(ta)
	printArea(sa)

}

// the getArea functions became field in the interface

func (t triangleArea) getArea() float64 {
	return t.height * t.base * 0.5
}

func (s squareArea) getArea() float64 {
	return s.sideLength * s.sideLength
}

// this function is used to implement the interface shape
func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

// func printArea(ta triangleArea) {
// 	fmt.Println(ta.getArea)
// }

// func printArea(sa squareArea) {
// 	fmt.Println(sa.getArea)
// }
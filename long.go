package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) getName() {
	fmt.Println(p.firstName)
}

func main() {
	longs := Person{firstName: "Longs", lastName: "Pemun", age: 30}
	longs.getName()
}

package main

import "fmt"

// embedding structs into struct fields
type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo //embed contactInfo structs here
}

// 1 create a new person struct
func main() {
	// // create a new person variable
	// oyin := person{"Oyindamola", "Abiola"} // firstName = "Oyindamola", lastName = "Abiola"
	// fmt.Println(oyin)

	// // create a new person variable
	// dassy := person{firstName: "Dasola", lastName: "Bamidele"} // firstName = "Dasola", lastName = "Bamidele"
	// fmt.Println(dassy) // prints the values of the variables only
	// // below prints with the key properties of the variable with \n (nextlint) atttribute
	// fmt.Printf("%+v\n", dassy) 

	// // create a new person variable
	// var efun person

	// // update the efun variable of type person
	// efun.firstName = "Efunroye"
	// efun.lastName = "Abosede"

	// fmt.Println(efun)
	// fmt.Printf("%+v\n", efun)


	// NEW STRUCT FUNCTION

	oyinda := person{
		firstName: "Oyindamola",
		lastName: "Abiola",
		contactInfo: contactInfo {
			email: "aodasola95@gmail.com",
			zipCode: 1234,
		},
	}

	// fmt.Printf("%+v\n", oyinda)

	// use the reciever function to update person (pointer shortcut)
	// oyinda.updatePerson("Arike")

	// to use the reciever function printPerson
	oyinda.printPerson()

	// pointer
	// oyindaPointer gives the object/struct (oyinda person) memory address...
	// ...because of the ampersand &oyinda
	oyindaPointer := &oyinda
	oyindaPointer.updatePerson("OyindaArike")
}

// pointer function to person
// * give access to the to the person object/struct (oyinda) itself and not the address this time around
// so *pointToPerson represents the object person itself(oyinda) because of the asteric and not just...
// ...the address like oyindaPointer 
func (pointerToPerson *person) updatePerson( newName string ) {
	(*pointerToPerson).firstName = newName
}

// simple pointer shortcut
// func (p *person) updatePerson( newName string ) {
// 	p.firstName = newName
// }

// a reciever function for person to print person 
// p here represents the object 'person'
func (p person) printPerson() {
	fmt.Printf("%+v", p)
}
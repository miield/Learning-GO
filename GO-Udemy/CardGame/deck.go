package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
	The goal here is to create a card deck game
	The step by step is below
	*/

// 1. create a type deck that represents the deck

// this is a data-type created: array-slice of string type
// and it's a slice of string, but deck type
type deck []string

// 2. the function and logic to create and return the list of playing cards in arrays of string, 
// basically returns a deck object
func newDeck() deck {
	/**
	 * to create the actual deck cards, we will have to spell
	 * the names long,
	 * so instead of that, we will create a nested for-loop to
	 * join the suits and values
	 */

	/**
	 * This syntax creates a new instance of the deck type by using the
	 * curly braces {}
	 * to enclose the values or elements that make up the new instance. In this case,
	 * the curly braces are empty, indicating that the cards variable is being initialized
	 * as an empty instance of the deck type.
	 */
	// cards is assigned to type deck
	cards := deck{}

	// Initialize the variables for the 2 properties of the deck, that are of slice pf string type
	cardSuits := []string{"Spades", "Heart", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// loop through both properties of the deck
	// i and j are indexes of the items in the cardSuits array and cardValues array
	// for i, suit := range cardSuits { // replace i with _ to stop the error of the unused var i
	for _, suit := range cardSuits {
		// for j, suit := range cardSuits { // replace j with _ to stop the error of the unused var j
		for _, value := range cardValues {
			// append is to add the items to/inside the cards slice of string also a deck type
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// 3. A receiver function to log out content the deck of playing cards

/** receiver function: stands like the methods associated/bonded with the type created in this case type deck */
/** d stands like a variable of type deck, it can also mean "this" in OOP
 * 	print() is the function name
 *  it is to print the list of cards in the deck and can be used to print a deck anywhere in the program
 */

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// 5. to create hands of cards
func deal(d deck, handSize int) (deck, deck) {
	// this means starts from the beginning of the slice, stops at int handsize
	// and starts from the end of the first handsize and stops at the end of the slice i.e int handsize
	return d[:handSize], d[handSize:]
	classPosition := []int}

// 6. save a list of cards to a file on the drive

// convert slice to a single string
// we have a receiver function here because we need to access the d of deck instance
// receiver function here means that we can only use it with the deck type
func (d deck) toString() string {
	// convert deck to slice of string
	// []string(d) // ...HERE
	// convert slice of string to string with the package strings
	return strings.Join([]string(d), ",")
}

// below is a save to file method so to be useful and called at anytime needed
// the para "filename" is made dynamic so that users can use their preferred filename
// returns error in case something goes wrong
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// this doesn't need a receiver function cos, we're reading from the file(local storage)
// then we convert it to a deck, therefore cannot have a receiver function
// this function is used to call the stand  package ReadFile method
func newDeckFromFile(filename string) deck {
	// bs = byte slice also means a string(list of cards seperated by commas)
	// this returns the
	// the err means if anything goes wrong, it returns the error, value, else return "nil" as the value
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error", err)
		// so if an error is returned we can stop the process entirely using the OS package using the exit function
		// 0 means no error found, 1 indicates error was found
		os.Exit(1)
	}
	// this below is a type conversion to convert the byte slice to a single string
	// string(bs)
	// to seperate the single string with comma to return slice of string
	s := strings.Split(string(bs), ",")
	//  another type conversion to Deck
	return deck(s)
}

// to shuffle the card
func (d deck) shuffle() {
	// creating true randomness
	// first create the source of randomness
	source := rand.NewSource(time.Now().UnixNano()) // time package, unixNano time to generate the int64 number
	// then create the randomness using the source
	r := rand.New(source)
	// range here is keyword for forloop in GO
	for i := range d {
		// below is creating the random number rand r is a package in GO, it does the randomness
		// newPosition is the random index of d/deck
		newPosition := r.Intn(len(d) - 1)

		// below is the swapping of the normal index and the random index
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

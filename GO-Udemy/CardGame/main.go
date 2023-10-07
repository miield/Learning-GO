package main

// import "fmt"

func main() {

	// newDeck method from the deck.go is being assigned to the new variable cards
	// cards := newDeck()

	// cards.print()

	// calling the deal function from the deck.go.
	// this doesn't need imports because it's in the same package
	// below the 2 variables are declared and assigned to the deal's parameters
	// cards = hand, 5 = remaining cards
	// hand, remainingCards := deal(cards, 5)

	// we are using the print function becos, print can only output deck types only in the forloop implementation
	// hand.print()
	// remainingCards.print()

	// cards := deck{"Six of spades", callCard()}
	// cards = append(cards, "Kings card")

	// print function is in deck.go, it uses forloop to print the cards
	// cards.print()

	// type conversion/typeastion
	// myName := "Oyindamola, unapologetic ambitious lady"
	// fmt.Println([]byte(myName))

	// converts card to string
	// fmt.Println(cards.toString())

	// save cards to file
	// cards.saveToFile("myCards")

	// read/load from the file
	// cards := newDeckFromFile("myCards")
	// print() still works with decks
	// cards.print()

	// to call the shuffled cards
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func callCard() string {
// 	return "Ace of spade"
// }

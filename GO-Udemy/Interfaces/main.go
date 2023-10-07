package main

import "fmt"

// this exercise deals with 2 different types of functions
// ...but still have some implementation in common

/** we need
1. we need to tie-up the function into a type,
and we're using struct type. we need them empty though
2. we need actual bot-type variables
3. we need the struct type to be empty just for this case
4. we need function english bot and spanish bot for the actual 
logic to return the greetings
5. functions (2 different) to print out the greetings, one for each
*/

// this interface type right here deals with the both struct bots...
//... because they own the the getGreeting functions
type bot interface {
	getGreeting() string
}

// we need to tie-up the function into a type, and we're using struct type...
//... we need them empty though...
type englishBot struct{}
type spanishBot struct{}

// function that implements the bot interfaces
func main() {
	// reps the bot variables of type englishBot & spanishBot...
	//...and indirectly the satisfy the bot interface itself
	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)
}

// the functions with the logic to return the greetings,...
// ...tied up into a type : englishBot & spanishBot
func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola hola!!!!"
}

// noted that the code has been optimised

// function to print out the greetings
// this function expects the a particular variable/function (of any bot type; eng/spnh)...
//... as parameter to the printGreetings function

// correct: this function here takes the bot interface type,... 
// ...returns the greetings because the greeting bots are connected to the bot and...
// ...to the getGreeting functions
// b represents the englishBot and spanishBot that are connected to the getGreeting function
func printGreetings(b bot) {
	fmt.Println(b.getGreeting())
}

// Old versions of print greetings
// func printGreetings(eb englishBot) {
// 	fmt.Println(eb.getGreetings())
// }

// func printGreetings(sb spanishBot) {
// 	fmt.Println(sb.getSpanishGreetings())
// }

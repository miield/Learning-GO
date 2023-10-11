package main

import (
	"fmt"
	"net/http"
	"time"
)

// status checking for several websites
func main() {
	// list of sites := variables
	links := []string {
		"http://www.google.com",
		"http://www.stackoverflow.com",
		"http://www.facebook.com",
		"http://www.twitter.com",
		"http://www.amazon.com",

	}

	// create channel instance
	// channels are exactly the same as variables, and should be treated as one
	c := make(chan string)

	for _, link := range links {
		// without goroutines
		// checkLink(link)
		// with goroutines
		go checkLink(link, c)
	}

	// print out the message received from the channel to the main 
	// ...routine that is virtually created
	// XXX this code below is a blocking line XXX
	// fmt.Println(<- c)
	// fmt.Println(<- c)

	// loop through the links.... this works
	// for i := 0; i < len(links); i++ {
	// 	// fmt.Println(<-c)
	// }

	// loop for eternity timely
	for l := range c {
		go func () {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}
}

// Implementation....
// create a function that checks if the link is responding to traffic
// create an http Get request and return the response 
// pass in the c, the channel, chan as the type and string as the data 
// type it can only work with is string
func checkLink(link string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link)

	// check if the link is responding
	if err != nil {
		fmt.Println(link, "might be down")
		// pass message to the channel
		c <- link
		return
	} 
	fmt.Println(link, "is responding")
	// pass message to the channel
	c <- link
}

// go-routines cannot work without the channels because of the main-routine 
// will not be able to communicate with the children-channels
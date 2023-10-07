package main

import (
	"fmt"
	"net/http"
)

// status checking for several websites
func main() {
	// list of sites
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

	// print out the message from the channel
	fmt.Println(<- c)
}

// create a function that checks if the link is responding to traffic
// create an http Get request and return the response 
// pass in the c, the channel, chan as the type and string as the data 
// type it can only work with is string
func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	// check if the link is responding
	if err != nil {
		fmt.Println(link, "might be down")
		// pass message to the channel
		c <- "Website is down"
		return
	} 
	fmt.Println(link, "is responding")
	// pass message to the channel
	c <- "Website is up again"
}

// go-routines cannot work without the channels because of the main-routine 
// will not be able to communicate with the children-channels
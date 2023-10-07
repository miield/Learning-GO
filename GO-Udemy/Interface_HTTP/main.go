package main

import (
	"net/http"
	"os"
	"fmt"
	"io"
)

func main() {
	resp, err := http.Get("http://google.com")
	// log out the error if one occours
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// log out the response
	// this just prints out bunch of responses that are not neccessarity useful...
	// ...  and unorganised
	// fmt.Println(resp)

	// using the reader interface to get the response body from google.com
	// the bs is created empty with 99999 size capacity, to be passed to the...
	// ...Read function to get the...response body
	// bs := make([]byte, 99999)
	// // read the google.com response body
	// resp.Body.Read(bs)
	// // converts the bytes into strings
	// fmt.Println(string(bs))


	// advance way of readng and getting the response body	
	io.Copy(os.Stdout, resp.Body)
}

/**
	When the Read method inside the Reader interface is called, it creates an empty
	byte slice and returns to us the number of data bytes that is being pushes into it, as the 
	response body. In a classic programming language, we suppose to get the slice
	of bytes back as the data requested straight up, not to first creat an empty slice where the
	response body is being pushed like the case of GO
*/

/**
	the writer interface implements the write function, which does the opposites of
	read methods. the writer interface takes the bytes of slice and sends it 
	out through an outputs channel e.g HTTP request, terminal etc.
    
	Copy function implementes the Writer Interface and the Reader Interface and returns int64 & error
	io.Copy above has 2 properties os.Stdout & resp.Body. The os.Stdout implement writer interface 
	practically send the bytes of slice
	os.Stdout -> File type -> Write() -> so the File type implements Write interface
*/

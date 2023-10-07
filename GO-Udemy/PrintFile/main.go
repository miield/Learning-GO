package main

import (
	"fmt"
	"os"
	"path/filepath"
	"io"
)

func main() {

	/** 
		1. Create a new file called file.txt, 
		2. write the code in line 19
		3. run "go run main.go file.txt"
		4. It brings up the file in the terminal
	*/

	// to get access to the file name from the command line arguments list
	// fmt.Println(os.Args[1])

	// EXERCISE: file
	// to open up the file on a hardrive matching the file name
	// to open the file, use the Open function provided by the os package
	// provide the file name as the argument and it will return a pointer 
	// to a File struct and error if it fails

	// exploring the File struct, check the interface that File struct implements
	// the File type implements the Reader interface

	// f reps the File struct, err reps the error
	f, err := os.Open(os.Args[1])

	// error handling 
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Open package implements File interface
	// the File is the f that implements the Read interface
	// the Copy implements the Writer interface which is the os.Stdout & f which implements
	// the read function 
	io.Copy(os.Stdout, f)
}

// to print out or to get access to the list of arguments that are passed 
// ... into the program we do
// fmt.Println(os.Args[0]) // os.Args[0] is a type of slice of strings

// note that there is an executable file created, saved in temporary dir 
// when the program is run, which takes the slot 0, i.e the first argument
// so to get the actual file path that we created, we do .... fmt.Println(os.Args[1])



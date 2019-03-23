package main

import (
	"os"
)

// If the error given is not nil, the program will be aborted
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Opening the file:
	input := "problems.csv"
	f, err := os.Open(input)
	check(err)      // Checking for errors
	defer f.Close() // Making the file close at the end of the execution
}

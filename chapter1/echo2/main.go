// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for i, arg := range os.Args[0:] { // here _ is empty(place holder) and this will produce only the values without the index
		s += sep + fmt.Sprintf("%v:%s", i, arg) // string print format %v:(String random value) %s(just string) %q(string and will put quotes around the string)
		sep = " "
	}
	fmt.Println(s)
}

//!- add an argument after the run command that be printed for example `go run echo1/main.go test` this will print test test1 test2
// Task1 Modify the echo program to also print os.Args[0], the name of the command that invoked it.

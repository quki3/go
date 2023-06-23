// package - collection of software file.
// main -special package named in go this must have a main() witch serves as a entri poin of the programs and It is used for creating standalone ejecutable programs
package main

// import - when you need your code to do something that might have been implemented by someone else
import (
	// fmt - contains functions for formatting text
	// rsc.io/quote - this package collects pithy sayings
	"fmt"
	"rsc.io/quote"
	)

// func - define a function
// main - create a enter point for go
func main() {
	//. Println - used for print and format ouput
	// quote - this is a package name of the collection of sofware file that come on rsc.io/quote
	// .Go - this is a function writed inside of quote package
	fmt.Println(quote.Go())
}

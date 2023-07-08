package main

import (
	"fmt"
	"os"
)

func main () {
	s := ""
	for i, arg := range os.Args[1:]{
		s +=arg
		fmt.Println(arg)
		fmt.Println( i, s)
		s=""
	}
}

package main

import (
	"fmt"
)

func main() {
	n1, n2 := Names() //return two string
	fmt.Println(n1,n2)

	n3, _ := Names()
	fmt.Println(n3)
}

func Names() ( first string, second string ) {
	first = "first string"
	second = "second string"
	return
}

package main

import "fmt"


var n int = 8
func main(){
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Println(x)
	}
}

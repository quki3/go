package main
import "fmt"

var x string = "lexical"
func main () {
	x := "local"
	if x == "lexical" {
		fmt.Printf("%c",x) // $ "local"
	} else if x == "local" {
		fmt.Println(x)
	} else {
		fmt.Println(x,x)
	}
	fmt.Println("x = ",x)
}

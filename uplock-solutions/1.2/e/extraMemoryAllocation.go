package main
import "fmt"
var reachable *int
var unreachable *int

func main(){
	f()
	g()
	fmt.Println(reachable,unreachable)
}

func f(){
	var x int
	x = 1
	reachable = &x
}
func g(){
	y := new(int)
	*y = 1
	//unreachable = &y // error $ cannot use &y (value of type **int) as *int value in assignment
}

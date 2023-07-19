package main
import "fmt"

func main(){
	fmt.Println(gcd(10,5))
}

func gcd(x, y int) int{
	for y != 0 {
		x, y = y, x%y
	}
	return x
}


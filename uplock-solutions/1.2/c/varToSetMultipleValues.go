package main
import (
	"fmt"
	"os"
)
func main(){
	var f, err = os.Open("test.txt")
	fmt.Println(f,err)
}

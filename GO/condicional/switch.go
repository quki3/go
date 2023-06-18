package main

import (
	"fmt"
)

func main(){
	switch {
	case 2==4:
		fmt.println("not must print")
	case 3==3:
		fmt.println("must print")
	case 4==5:
		fmt.println("not must print")
	default:
		fmt.println("default print")
	}
}

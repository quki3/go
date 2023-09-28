package main 

import (
	"fmt"
	"time"
)

func printCount ( c chan int ) {
	num := 0
	fmt.Println(num,"#1")
	for num >= 0 {
		fmt.Println(num,"#1.9 before num <- c")
		num = <- c			// num  wait for channel c send 

		fmt.Println(num,"#2") 		// way not print in console?
						// this case is because channel wait for
						// date to proccess
	}
}

func main() {
	c := make(chan int) 			// Interger-type channel c is created 
						// to communocate between goroutines
						// the channel needs to be identified as an interger
						// channel
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}

	go printCount(c)			// started as a goroutine
	fmt.Println(c,"#3")

	for _, v := range a {			// Liist of intergers is interated over and passed into
						// the channel c one at a time
		c <-v
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}

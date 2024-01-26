package main

import "fmt"

func main() {

    if 7%2 == 0 {
        fmt.Println("7 is even") // even balance uniform, etc
    } else {
        fmt.Println("7 is odd")// odd is strengh
    }

    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    if num := 9; num < 0 { //declare a variable here 
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

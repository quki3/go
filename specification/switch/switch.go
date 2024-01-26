package main

import (
    "fmt"
    "time"
)

func main() {

    i := 2 				//declaration
    fmt.Print("Write ", i, " as ")	// print
    switch i {				// declare the switch
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    i2 = 4
    fmt.Print("write", i2, "as")
    switch i2 {
    case 3:
	    fmt.Print("three")
    case 4:
	    fmt.Print("four")
    }

    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {  	//return the type
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}

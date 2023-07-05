package main

import "fmt"

func printValue(val interface{}) {
    fmt.Println("Value:", val)
}

func main() {
    var x interface{} // x is nil and has static type interface{}
    
    x = 42        // Assigning an int
    printValue(x) // Prints: Value: 42
    
    x = "Hello"   // Assigning a string
    printValue(x) // Prints: Value: Hello
    
    x = true      // Assigning a boolean
    printValue(x) // Prints: Value: true
}

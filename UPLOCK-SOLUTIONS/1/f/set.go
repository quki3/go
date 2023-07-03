package main

import "fmt"

// Define a struct type
type Person struct {
    name string
    age  int
}

// Define a method associated with the Person struct
func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
    return p.name
}

func main() {
    // Create an instance of the Person object
    person := Person{
        name: "John",
        age:  30,
    }

    // Call the SayHello method on the person object
    person.SayHello()
}

package main

import "fmt"

// Define a struct type
type Person struct {
    name string
    age  int
}
// define a struct dog
type Dog struct {
    name string
    age int
}

// Define a method associated with the Person struct
func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
    return p.name
}
func (d Dog) SayGuau(){
    fmt.Printf("guau, guau")
    return d.name
    }

func main() {
    // Create an instance of the Person object
    person := Person{
        name: "John",
        age:  30,
    }
    dog := Dog{
        name:"till",
        age:2,
        }

    // Call the SayHello method on the person object
    person.SayHello()
    dog.SayGuau()
}

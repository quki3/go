package main
import "fmt"

func main (){
  s := "hello, World"
  fmt.Println(len(s))
  fmt.Printlm(s[0], s[7])
  c := s[len(s))]  // panic: index out of range
  fmt.Println(s[0:5])
  fmt.Println(s[:5]) //"hello"
  fmt.Println(s[7:]) //"world"
  fmt.Println(s[:]) //"hello, world"
  fmt.Println("goodbye"+ s[5:]) //"goodbye, world"
}

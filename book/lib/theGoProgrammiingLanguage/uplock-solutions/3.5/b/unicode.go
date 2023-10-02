package main
import "fmt"

func main(){
  for i, r := range "Hello, BF"{
    fmt.Printf("%d\t%q\t%[1]d\n",i,r)
}

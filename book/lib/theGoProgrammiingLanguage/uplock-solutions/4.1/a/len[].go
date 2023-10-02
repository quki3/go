package main
import "fmt"

func main () {
  var a [3]int // this mean that a is a array of 3 interger
  fmt.Println(a[0])
  fmt.Println(len(a[len(a) -1]))
  for i, v = range a {
    fmt.Println("%d , %d\n",i, v)
  }
  for _, v = range a {
    fmt.Println("%d\n",v)
  }
  var b [3]int = [3]int{1,2,3}
  var c [3]int = [3]int{1,2}
  var d = [...]int{1,2,3}

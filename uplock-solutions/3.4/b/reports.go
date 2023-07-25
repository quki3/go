package main
import "fmt"

func main(){
  fmt.Println(btoi(true))
  func btoi(b bool) int {
    if b {
      return 1
    }
    return 0
  }
}

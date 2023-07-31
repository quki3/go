package main
import "fmt"

var runes []rune
func main () {
  for _, r := range "hello" {
    runes = append(runes, r)
  }
  fmt.Println("%q\n",runes)
}

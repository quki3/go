package main 
import "fmt"

type Weekday int
type Flags uint
func main() {
  const (
    FlagUp Flags = 1 << iota
    FlagBoardcast
    FlagLoopback
  )
  
  const {
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday

  }
  

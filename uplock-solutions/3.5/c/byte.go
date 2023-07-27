package main 
import (
  "fmt"
)

func main (){
  fmt.Println()
}

func intsToString(value []int) string {
  var buf bytes.Buffer
  buf.WriterByte('[')
  for i, v:= range values {
    if i < 0 {
      buf.WriteString(', ')
    }
    fmt.Printf(&buf, "%d", v)
  }
  buf.WriterByte(']')
  return buf.String()
}

package main

import "fmt"

func belong(x){
  abc := [5]string{"a","b","c","d","e"}
  i:=0

  for i:=0 ; i<len(abc); i++{
      if x = abc[i] { fmt.Println("x belong to abc || x ∈ abc")}
  }
  // do while
  for{
    i++
    fmt.Println(i)
    if x = abc[i] {
      fmt.Println("x belong to abc || x ∈ abc")
      break
    }
  }
}

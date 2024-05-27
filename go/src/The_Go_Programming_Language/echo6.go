package main

import (
  "fmt"
  "os"
)

func main(){
  var value string
  for i := 1; i < len(os.Args); i++ {
    value = os.Args[i]
    fmt.Println(i, ": ", value)
  }
}

// this is the while loop in C

package main

import "fmt"

func main(){
  sum := 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}

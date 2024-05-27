package main

import (
	"fmt"
)

func main() {
	x := factorial(20)
	fmt.Println(x)
}

func factorial (x int) int {
	if x == 0 {
		return 1
	} else {
		return factorial(x - 1) * x
	}
}

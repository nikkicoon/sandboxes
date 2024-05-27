package main

import (
	"fmt"
)

func main() {
	i := 7
	fmt.Println("Value of i: ", i)
	fmt.Println("i points to Memory Address: ", &i)
	x := 7
	fmt.Println(x, &x)
	inc(&x)
	fmt.Println(x, &x)
}

// Accept pointer to int to access and modify the original version.
func inc(x *int) {
	// dereference the pointer,
	// without the asterisk we would increment the memory address
	*x++
}

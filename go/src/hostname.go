package main

import (
	"fmt"
	"os"
)

func main() {
	n, _ := os.Hostname()
	fmt.Println(n)
}

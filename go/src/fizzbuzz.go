package main

import "fmt"

func main() {
	for i := 0; i <= 42; i += 2 {
		if (i%6 == 0) && (i%7 == 0) {
			fmt.Println("Fizz Buzz")
		} else if i%6 == 0 {
			fmt.Println("Fizz")
		} else if i%7 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

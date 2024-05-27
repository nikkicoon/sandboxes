package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {
	result := sum(3, 4)
	resultSqrt, errSqrt := sqrt(12.00)
	fmt.Println(result)
	if errSqrt != nil {
		fmt.Println(errSqrt)
	} else {
		fmt.Println(resultSqrt)
	}
}

// is this the same as x int, y int
func sum(x, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

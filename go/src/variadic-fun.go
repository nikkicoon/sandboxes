package main

import "fmt"

/*
the unpack operator ends with ... like slice...
while the pack operator starts with ... like ...Type
*/

/*
getMultiples function:
first argument is 'factor' of type 'int', which is a factor
of multiplication and later variable arguments ('variadic arguments')
of type 'int' are packed into the slice 'args'.
*/
func getMultiples(factor int, args ...int) []int {
	// create an empty slice with make, with length equal to the length
	// of args which is a slice.
	multiples := make([]int, len(args))

	// we are multiplying factor with elements of args and save them
	// in multiples.
	for index, val := range args {
		multiples[index] = val * factor
	}

	// return the slice multiples.
	return multiples
}

func main() {
	s := []int{10, 20, 30}
	mult1 := getMultiples(2, s...)
	mult2 := getMultiples(3, 1, 2, 3, 4)

	fmt.Println(mult1)
	fmt.Println(mult2)
}

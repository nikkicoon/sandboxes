// cf converts its numeric argument to Celsius, Fahrenheit, and Kelvin.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nikkicoon/go-tempconv/tempconv"
)

func main() {
	// range returns 2 values, index and copy of element
	// at that index. throw away the index by using '_'.
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
		fmt.Printf("%s = %s\n", c, tempconv.CToK(c))
		fmt.Printf("%s = %s\n", c, tempconv.CToRé(c))
		fmt.Printf("%s = %s\n", c, tempconv.CToRa(c))
		fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
		fmt.Printf("%s = %s\n", f, tempconv.FToK(f))
		fmt.Printf("%s = %s\n", f, tempconv.FToRé(f))
		fmt.Printf("%s = %s\n", f, tempconv.FToRa(f))
		fmt.Printf("%s = %s\n", k, tempconv.KToC(k))
		fmt.Printf("%s = %s\n", k, tempconv.KToF(k))
		fmt.Printf("%s = %s\n", k, tempconv.KToRé(k))
		fmt.Printf("%s = %s\n", k, tempconv.KToRa(k))
	}
}

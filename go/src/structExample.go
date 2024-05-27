package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main() {
	p := person{name: "Jake", age: 24}
	p2 := person{"Nini", 32}
	fmt.Println(p)
	fmt.Println(p.name)
	fmt.Println(p2, p2.name)
}

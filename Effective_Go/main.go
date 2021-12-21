package main

import (
	"fmt"
)

func main() {
	a := int(uint(0))  // largest int
	b := uint(^uint(0))  // largest int
	c := int(^uint(0) >> 1)  // largest int
	d := int(^uint(0) >> 1)  // largest int
	e := int(^uint(0) >> 1)  // largest int

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
}

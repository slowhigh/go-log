package main

import (
	"fmt"
)

func bufferedInt (bc chan int) {
	fmt.Println("bufferedInt")

	bc <- 10
	bc <- 20
}

func main () {
	bc := make(chan int, 2)

	bufferedInt(bc)

	fmt.Println(<- bc)
	fmt.Println(<- bc)
}

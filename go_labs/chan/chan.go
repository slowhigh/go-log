package main

import (
	"fmt"
)

func f1(v int, c3 chan chan chan int) {
	c2 := make(chan chan int)
	fmt.Println("f1	// c2 := make(chan chan int)")

	c3 <- c2
	fmt.Println("f1	// c3 <- c2")

	c1 := <-c2
	fmt.Println("f1	// c1 := <- c2")

	c1 <- v + 10
	fmt.Println("f1	// c1 <- v + 10")
}

func main() {
	c3 := make(chan chan chan int)
	fmt.Println("main	// c3 := make(chan chan chan int)")

	go f1(1, c3)
	fmt.Println("main	// go c1(1, c3)")

	c2 := <-c3
	fmt.Println("main	// c2 := <- c3")

	c1 := make(chan int)
	fmt.Println("main	// c1 := make(chan int)")

	c2 <- c1
	fmt.Println("main	// c2 <- c1")

	re := <-c1
	fmt.Println("main	// re := <- c1")

	fmt.Println("main	//", re)
}

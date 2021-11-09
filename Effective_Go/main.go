package main

import (
	"fmt"
	"time"
)

func b() {
	for i := 20; i < 30; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func a() {
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}

	time.Sleep(time.Second * 3)

	for i := 10; i < 20; i++ {
		fmt.Printf("%d ", i)
	}
}

func main() {
	var v  []int = make([]int, 5, 10) 
	v[0] = 1;
	v[1] = 1;
	v[2] = 1;
	v[3] = 1;
	v[4] = 1;
	v[5] = 1
	v[6] = 1;
	v[7] = 1;
	v[8] = 1;
	v[9] = 1;
	v[10] = 1;
}

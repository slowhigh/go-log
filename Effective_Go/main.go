package main

import (
	"fmt"
)

type ByteSize float64

const (
    _           = iota // iota = 0, ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota) // iota = 1
    MB // = 1 << (10 * iota) , iota = 2
    GB
    TB
    PB
    EB
    ZB
    YB
)

func main() {
	x := []int{1,2,3}
	y := []int{4,5,6}
	x = append(x, y...)
	fmt.Println(x)
}

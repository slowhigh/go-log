package main

import (
	"fmt"
)

type ByteSlice []byte

func (slice ByteSlice) Append1(data []byte) []byte {
    return append(slice, data...)
}

func (p *ByteSlice) Append2(data []byte) {
    slice := *p
    slice = append(slice, data...)
    *p = slice
}

func main() {
	var x ByteSlice = []byte{1, 2, 3}
	var y ByteSlice = []byte{4, 5, 6}
	
	fmt.Printf("append = %x \n", x.Append1(y))
	fmt.Printf("x = %x \n", x)
	fmt.Printf("y = %x \n", y)


    x.Append2(y)

    fmt.Printf("new x = %x", x)
}


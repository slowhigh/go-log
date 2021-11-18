package main

import (
	"fmt"
)

func Append(slice, data []byte) []byte {
	fmt.Printf("1-  %v - %p[%p]\n", slice, &slice, slice)
    l := len(slice)
    if l + len(data) > cap(slice) {  // reallocate
        // Allocate double what's needed, for future growth.
        newSlice := make([]byte, (l+len(data))*2)
		fmt.Printf("2- %v - %p[%p]\n", newSlice, &newSlice, newSlice)
        // The copy function is predeclared and works for any slice type.
        copy(newSlice, slice)
		fmt.Printf("3- %v - %p[%p]\n", newSlice, &newSlice, newSlice)
        slice = newSlice
    } 
	fmt.Printf("4- %v - %p[%p]\n", slice, &slice, slice)
    slice = slice[0:l+len(data)]
	fmt.Printf("5- %v - %p[%p]\n", slice, &slice, slice)
    copy(slice[l:], data)
	fmt.Printf("6- %v - %p[%p]\n", slice, &slice, slice)
    return slice
}

func main() {
	slice := make([]byte, 5, 10)
	slice2 := []byte { 2,2,2,2,2,2 }
	Append(slice, slice2)
}

// 1- [0 0 0 0 0] - 0xc000004078[0xc0000120b0]
// 2- [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] - 0xc0000040c0[0xc000014198]
// 3- [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] - 0xc0000040c0[0xc000014198]
// 4- [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] - 0xc000004078[0xc000014198]
// 5- [0 0 0 0 0 0 0 0 0 0 0] - 0xc000004078[0xc000014198]
// 6- [0 0 0 0 0 2 2 2 2 2 2] - 0xc000004078[0xc000014198]

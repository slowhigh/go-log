package main

import (
	"fmt"
)

func main() {
	// Allocate the top-level slice, the same as before.
	picture := make([][]uint8, 5) // One row per unit of y.
	// Allocate one large slice to hold all the pixels.
	pixels := make([]uint8, 3*5) // Has type []uint8 even though picture is [][]uint8.
	// Loop over the rows, slicing each row from the front of the remaining pixels slice.
	for i := range picture {
		picture[i], pixels = pixels[:3], pixels[3:]
        fmt.Printf("- %v - %p[%p]\n", picture, &picture, picture)
        fmt.Printf("- %v - %p[%p]\n", picture[i], &picture[i], picture[i])
		fmt.Printf("- %v - %p[%p]\n", pixels, &pixels, pixels)
	}
}

// 1- [0 0 0 0 0] - 0xc000004078[0xc0000120b0]
// 2- [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] - 0xc0000040c0[0xc000014198]
// 3- [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] - 0xc0000040c0[0xc000014198]
// 4- [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] - 0xc000004078[0xc000014198]
// 5- [0 0 0 0 0 0 0 0 0 0 0] - 0xc000004078[0xc000014198]
// 6- [0 0 0 0 0 2 2 2 2 2 2] - 0xc000004078[0xc000014198]

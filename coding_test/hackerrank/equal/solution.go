package equal

import (
	"fmt"
	"math"
	"sort"
)

// Equal
// https://www.hackerrank.com/challenges/equal/problem?isFullScreen=false

type cast struct {
	count int32
	piece int32
	arr   []int32
}

func equal(arr []int32) int32 {
	pieces := []int32{1, 2, 5}

	min := int32(math.MaxInt32)
	stack := make([]cast, 0)

	sort.Slice(arr, func(i, j int) bool { return arr[i] > arr[j] })

	//fmt.Printf("sort %+v\n", arr)

	for _, piece := range pieces {
		if arr[0] >= piece {
			newArr := make([]int32, len(arr))
			copy(newArr, arr)
			stack = append(stack, cast{1, piece, newArr})
		}
	}

	fmt.Printf("stack %+v\n", stack)

	// for len(stack) != 0 {
        for i := 0; i < 2; i++ {
		c := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		//fmt.Printf("dequeue %+v\n", c)
		fmt.Printf("stack %+v\n", stack)
		c.arr[0] -= c.piece

		isSame := true
		for i := 0; i < len(c.arr)-1; i++ {
			if c.arr[i] != c.arr[i+1] {
				isSame = false
			}
		}

		if isSame {
			//fmt.Printf("same %+v\n", c)
			if c.count < min {
				min = c.count
			}
			continue
		}

		sort.Slice(c.arr, func(i, j int) bool { return c.arr[i] > c.arr[j] })
		//fmt.Printf("sort %+v\n", c.arr)

		for _, piece := range pieces {
			if c.arr[0] >= piece {
				newArr := make([]int32, len(c.arr))
				copy(newArr, c.arr)
				stack = append(stack, cast{c.count + 1, piece, newArr})
			}
		}

		//fmt.Printf("queue %+v\n", queue)
	}

	return min
}

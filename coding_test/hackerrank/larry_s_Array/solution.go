package larry_s_array

import (
	"errors"
	"fmt"
)

// Larry's Array
// https://www.hackerrank.com/challenges/larrys-array/problem?isFullScreen=false

type node struct {
	A     []int32
	Pivot int
}

func larrysArray(A []int32) string {
	var (
		res       = "NO"
		queue     = make([]node, 0)
		lastPivot = len(A) - 2
	)

	before := make([]int32, len(A))
	copy(before, A)
	for i := 0; i < 3; i++ {
		newA := make([]int32, len(A))
		before, _ = rotate(before, 0)
		copy(newA, before)
		queue = append(queue, node{A: newA, Pivot: 0})
	}

	fmt.Printf("%+v\n", queue)

	for len(queue) > 0 {
		target := queue[0]
		queue = queue[1:]
		fmt.Printf("%+v\n", queue)

		if isSorted(target.A) {
			res = "YES"
			break
		}

		if target.Pivot + 1 >= lastPivot {
			continue
		}

		before := make([]int32, len(target.A))
		copy(before, target.A)
		for i := 0; i < 3; i++ {
			newA := make([]int32, len(target.A))
			before, _ = rotate(before, target.Pivot + 1)
			copy(newA, before)
			queue = append(queue, node{A: newA, Pivot: target.Pivot + 1})
		}
		fmt.Printf("%+v\n", queue)
	}
	fmt.Printf("%+v\n\n\n", queue)

	return res
}

func isSorted(arr []int32) bool {
	len := int32(len(arr))

	for i := int32(0); i < len; i++ {
		if i+1 != arr[i] {
			return false
		}
	}

	return true
}

func rotate(arr []int32, pivot int) ([]int32, error) {
	if len(arr)-2 <= pivot {
		return nil, errors.New("invalid pivot")
	}

	front := arr[:pivot]
	target := []int32{arr[pivot+1], arr[pivot+2], arr[pivot]}
	back := arr[pivot+3:]

	rot := append(append(front, target...), back...)

	return rot, nil
}

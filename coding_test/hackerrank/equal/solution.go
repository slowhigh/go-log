package equal

import (
	"math"
)

// Equal
// https://www.hackerrank.com/challenges/equal/problem?isFullScreen=false

func equal(arr []int32) int32 {
	var (
		min, max int32
		pieces   = []int32{5, 2, 1}
		count    = int32(0)
		pivot    = 0
	)

	for {
		min, max = int32(math.MaxInt32), int32(0)
		for i, v := range arr {
			if v < min {
				min = v
			}

			if max < v {
				max = v
				pivot = i
			}
		}

		for _, v := range pieces {
			if min <= arr[pivot]-v {
				arr[pivot] = arr[pivot] - v
				count++
				continue
			}
		}

		equal := true
		for i := 1; i < len(arr); i++ {
			if arr[i-1] != arr[i] {
				equal = false
				break
			}
		}

		if equal {
			break
		}
	}

	return count
}

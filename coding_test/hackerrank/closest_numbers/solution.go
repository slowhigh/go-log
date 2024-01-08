package closest_numbers

import (
	"fmt"
	"sort"
)

// Closest Numbers
// https://www.hackerrank.com/challenges/closest-numbers/problem?isFullScreen=false

func closestNumbers(arr []int32) []int32 {
	var (
		minDiff = int64(-1)
		pairArr = make([]int32, 0)
	)

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	fmt.Printf("sorted %+v\n", arr)

	for i := 1; i < len(arr); i++ {
		diff := int64(arr[i]) - int64(arr[i-1])
		fmt.Printf("%d - %d = %d\n", arr[i], arr[i-1], minDiff)

		if diff < 0 {
			diff *= -1
		}

		if minDiff == -1 || diff < minDiff {
			minDiff = diff
		}
	}

	fmt.Printf("min diff -> %d", minDiff)

	for i := 1; i < len(arr); i++ {
		diff := int64(arr[i]) - int64(arr[i-1])

		if diff == minDiff || diff == minDiff*-1 {
			pairArr = append(pairArr, arr[i-1], arr[i])
		}
	}

	return pairArr
}

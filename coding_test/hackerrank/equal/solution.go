package equal

import (
	"math"
)

// Equal
// https://www.hackerrank.com/challenges/equal/problem?isFullScreen=false
func equal(arr []int32) int32 {
	var (
		min   = int32(math.MaxInt32)
		count = int32(0)
		r     int32
	)

	for _, v := range arr {
		if v < min {
			min = v
		}
	}

	for _, v := range arr {
		v -= min

		if v >= 5 {
			r = v / 5
			v -= 5 * r
			count += r
		}

		if v >= 2 {
			r = v / 2
			v -= 2 * r
			count += r
		}

		if v >= 1 {
			r = v / 1
			v -= 1 * r
			count += r
		}
	}

	return count
}

package expected_table_of_matches

import (
	"math"
)

func solution(n int, a int, b int) int {
	depth := int(math.Log2(float64(n)))
	left, right := 1, n
	var center int

	for {
		center = ((right - left - 1) / 2) + left
		if (a <= center && center < b) || (b <= center && center < a) {
			break
		} else if a <= center && b <= center {
			right = center
			depth--
		} else if center < a && center < b {
			left = center + 1
			depth--
		}
	}

	return depth
}

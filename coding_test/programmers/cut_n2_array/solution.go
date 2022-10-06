package cut_n2_array

import "math"

func solution(n int, left int64, right int64) []int {
	len := right - left + 1
	result := make([]int, len)

	for i := int64(0); i < len; i++ {
		row := ((left + i) / int64(n)) + 1
		col := ((left + i) % int64(n)) + 1

		result[i] = int(math.Max(float64(row), float64(col)))
	}

	return result
}

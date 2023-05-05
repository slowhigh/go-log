package count_and_addition_of_divisor

import (
	"math"
)

func solution(left int, right int) int {
	if left == 1 && right == 1 {
		return -1
	}

	sum := 0

	for i := left; i <= right; i++ {
		sqrt := math.Sqrt(float64(i))

		if math.Mod(sqrt, 1) == 0 {
			sum -= i
		} else {
			sum += i
		}
	}

	return sum
}

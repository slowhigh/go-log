package internal

func solution(a []int, b []int) int {
	sum := 0

	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}

	return sum
}

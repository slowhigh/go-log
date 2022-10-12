package fibonacci_number

func solution(n int) int {
	a, b := 0, 1
	var c int

	for i := 1; i < n; i++ {
		c = (a + b) % 1234567
		a, b = b, c
	}

	return c
}

package long_jump

func solution(n int) int64 {
	if n < 3 {
		return int64(n)
	}

	n1, n2 := int64(1), int64(2)
	n3 := int64(0)

	for i := 2; i < n; i++ {
		n3 = (n1 + n2) % 1234567
		n1, n2 = n2, n3
	}

	return n3
}

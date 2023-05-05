package expected_table_of_matches

func solution(n int, a int, b int) int {
	count := 1

	for {
		if a%2 == 1 {
			a++
		}
		if b%2 == 1 {
			b++
		}
		if a == b {
			break
		}
		count++
		a /= 2
		b /= 2
	}

	return count
}

package manasaAndStones

// Manasa and Stones
// https://www.hackerrank.com/challenges/manasa-and-stones/problem?isFullScreen=false

func stones(n int32, a int32, b int32) []int32 {
	check := make(map[int32]bool)
	lastStones := make([]int32, 0)

	if a > b {
		a, b = b, a
	}

	for i := int32(0); i < n; i++ {
		lastStone := (b * i) + (a * (n - 1 - i))

		if ok := check[lastStone]; !ok {
			check[lastStone] = true
			lastStones = append(lastStones, lastStone)
		}
	}

	return lastStones
}

package countingValleys

// Counting Valleys
// https://www.hackerrank.com/challenges/counting-valleys/problem?isFullScreen=false

func countingValleys(steps int32, path string) int32 {
	hight, u, d, count := int32(0), rune('U'), rune('D'), int32(0)

	for _, r := range path {
		if r == u {
			if hight == -1 {
				count++
			}

			hight++
		} else if r == d {
			hight--
		}
	}

	return count
}

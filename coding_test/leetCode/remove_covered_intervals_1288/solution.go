package remove_covered_intervals_1288

func removeCoveredIntervals(intervals [][]int) int {
	count := 0
	intervalLen := len(intervals)
	skipArr := make([]bool, intervalLen)

	for i := 0; i < intervalLen; i++ {
		if skipArr[i] {
			continue
		}

		isUniqued := true

		for k := i + 1; k < intervalLen; k++ {
			target := intervals[i]
			next := intervals[k]
			if next[0] <= target[0] && target[1] <= next[1] {
				isUniqued = false
				break
			}

			if target[0] <= next[0] && next[1] <= target[1] {
				skipArr[k] = true
			}
		}

		if isUniqued {
			count++
		}
	}

	return count
}
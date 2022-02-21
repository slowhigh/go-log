package main

// Given an array intervals where intervals[i] = [li, ri] represent the interval [li, ri), remove all intervals that are covered by another interval in the list.

// The interval [a, b) is covered by the interval [c, d) if and only if c <= a and b <= d.

// Return the number of remaining intervals.

// Example 1:
// Input: intervals = [[1,4],[3,6],[2,8]]
// Output: 2
// Explanation: Interval [3,6] is covered by [2,8], therefore it is removed.

// Example 2:
// Input: intervals = [[1,4],[2,3]]
// Output: 1

// Constraints:

// 1 <= intervals.length <= 1000
// intervals[i].length == 2
// 0 <= li <= ri <= 105
// All the given intervals are unique.

import (
	"fmt"
)

func removeCoveredIntervals(intervals [][]int) int {
	count := 0
	intervalLen := len(intervals)
	skipArr := make([]bool, intervalLen)

	for i := 0; i < intervalLen; i++ {
		if skipArr[i] {
			continue
		}
		
		isUniqued := true

		for k := i+1; k < intervalLen; k++ {
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

func main() {
	intervals := [][]int {{1,4},{3,6},{2,8}}
	intervals2 := [][]int {{1,4},{2,3}}

	fmt.Println(removeCoveredIntervals(intervals))
	fmt.Println(removeCoveredIntervals(intervals2))
}
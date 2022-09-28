package h_index

import "sort"

// citations: []int {3, 0, 6, 1, 5},
// result: 3,

// 6 5 3 1 0
// 0 <= h <= n
// 1 2 3 4 5

func solution(citations []int) int {
	sort.Ints(citations)
	len := len(citations)

	for i := 1; i <= len; i++ {
		if citations[len-i] < i {
			return i - 1
		}
	}

	return len
}

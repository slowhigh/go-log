package h_index

import "sort"

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

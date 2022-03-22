package search_in_rotated_sorted_array_33

import "testing"

type testCase struct {
	nums   []int
	target int
	output int
}

func Test_search(t *testing.T) {
	testCaseArr := []testCase{
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			output: 4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			output: -1,
		},
		{
			nums:   []int{1},
			target: 0,
			output: -1,
		},
		{
			nums:   []int{1},
			target: 1,
			output: 0,
		},
		{
			nums:   []int{1, 3, 5},
			target: 5,
			output: 2,
		},
		{
			nums:   []int{5, 1, 3},
			target: 1,
			output: 1,
		},
	}

	for i, tc := range testCaseArr {
		if search(tc.nums, tc.target) != tc.output {
			t.Errorf("Wrong output - %d", i)
		}
	}
}

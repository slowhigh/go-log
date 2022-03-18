package find_minimum_in_rotated_sorted_array_153

import "testing"

type testCase struct {
	nums   []int
	output int
}

func Test_findMin(t *testing.T) {
	testCaseArr := []testCase {
		{
			nums: []int {3,4,5,1,2},
			output: 1,
		},
		{
			nums: []int {4,5,6,7,0,1,2},
			output: 0,
		},
		{
			nums: []int {11,13,15,17},
			output: 11,
		},
	}

	for i, tc := range testCaseArr {
		if findMin(tc.nums) != tc.output {
			t.Errorf("Wrong output - %d", i)
		}
	}
}

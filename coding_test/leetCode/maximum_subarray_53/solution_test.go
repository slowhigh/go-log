package maximum_subarray_53

import "testing"

type testCase struct {
	nums []int
	output int
}

func Test_maxSubArray(t *testing.T) {
	testCaseArr := []testCase {
		{
			nums: []int {-2,1,-3,4,-1,2,1,-5,4},
			output: 6,
		},
		{
			nums: []int {1},
			output: 1,
		},
		{
			nums: []int {5,4,-1,7,8},
			output: 23,
		},
	}

	for i, tc := range testCaseArr {
		if maxSubArray(tc.nums) != tc.output {
			t.Errorf("Wrong output - %d", i)
		}
	}
}
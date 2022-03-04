package subarray_sum_equals_k_560

import "testing"

type testCase struct {
	nums   []int
	k      int
	output int
}

func Test_subarraySum(t *testing.T) {
	testCaseArr := []testCase{
		{
			nums:   []int{1, 1, 1},
			k:      2,
			output: 2,
		},
		{
			nums:   []int{1, 2, 3},
			k:      3,
			output: 2,
		},
	}

	for _, tc := range testCaseArr {
		if subarraySum(tc.nums, tc.k) != tc.output {
			t.Error("Wrong output")
		}
	}
}

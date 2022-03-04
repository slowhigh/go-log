package majority_element_169

import "testing"

type testCase struct {
	nums   []int
	output int
}

func Test_majorityElement(t *testing.T) {
	testCaseArr := []testCase{
		{
			nums:   []int{3, 2, 3},
			output: 3,
		},
		{
			nums:   []int{2, 2, 1, 1, 1, 2, 2},
			output: 2,
		},
	}

	for _, tc := range testCaseArr {
		if tc.output != majorityElement(tc.nums) {
			t.Error("Wrong output")
		}
	}
}

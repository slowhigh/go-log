package single_number_136

import "testing"

type testCase struct {
	nums   []int
	output int
}

func Test_singleNumber(t *testing.T) {
	testCaseArr := []testCase{
		{
			nums:   []int{2, 2, 1},
			output: 1,
		},
		{
			nums:   []int{4, 1, 2, 1, 2},
			output: 4,
		},
		{
			nums:   []int{1},
			output: 1,
		},
	}

	for _, tc := range testCaseArr {
		if singleNumber(tc.nums) != tc.output {
			t.Error("Wrong output")
		}
	}
}

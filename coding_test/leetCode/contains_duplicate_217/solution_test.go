package contains_duplicate_217

import "testing"

type testCase struct {
	nums []int
	output bool
}

func Test_containsDuplicate(t *testing.T) {
	testCaseArr := []testCase {
		{
			nums: []int {1,2,3,1},
			output: true,
		},
		{
			nums: []int {1,2,3,4},
			output: false,
		},
		{
			nums: []int {1,1,1,3,3,4,3,2,4,2},
			output: true,
		},
	}

	for i, tc := range testCaseArr {
		if containsDuplicate(tc.nums) != tc.output {
			t.Errorf("Wrong output - %d", i)
		}
	}
}
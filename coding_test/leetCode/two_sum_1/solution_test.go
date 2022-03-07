package two_sum_1

import (
	"testing"
	"reflect"
)

type testCase struct {
	nums []int
	target int
	output []int
}

func Test_twoSum(t *testing.T) {
	testCaseArr := []testCase {
		{
			nums: []int {2,7,11,15},
			target: 9,
			output: []int {0,1},
		},
		{
			nums: []int {3,2,4},
			target: 6,
			output: []int {1,2},
		},
		{
			nums: []int {3,3},
			target: 6,
			output: []int {0, 1},
		},
	}

	for _, tc := range testCaseArr {
		if !reflect.DeepEqual(twoSum2(tc.nums, tc.target), tc.output) {
			t.Error("Wrong output")
		}
	}
}

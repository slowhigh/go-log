package three_sum_15

import (
	"reflect"
	"testing"
)

type testCase struct {
	nums []int
	output [][]int
}

func Test_threeSum(t *testing.T) {
	testCaseArr := []testCase {
		{
			nums: []int {-1,0,1,2,-1,-4},
			output: [][]int {{-1,-1,2},{-1,0,1}},
		},
		{
			nums: []int {},
			output: [][]int {},
		},
		{
			nums: []int {0},
			output: [][]int {},
		},
	}

	for i, tc := range testCaseArr {
		if !reflect.DeepEqual(threeSum(tc.nums), tc.output) {
			t.Errorf("Wrong output - %d", i)
		}
	}
}
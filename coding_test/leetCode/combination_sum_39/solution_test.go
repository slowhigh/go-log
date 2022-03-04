package combination_sum_39

import (
	"reflect"
	"testing"
)

type testCase struct {
	candidates []int
	target     int
	output     [][]int
}

func Test_combinationSum(t *testing.T) {
	testCaseArr := []testCase{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			output:     [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates: []int{2, 3, 5},
			target:     8,
			output:     [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			candidates: []int{2},
			target:     1,
			output:     [][]int{},
		},
	}

	for _, tc := range testCaseArr {
		output := combinationSum(tc.candidates, tc.target)
		if !reflect.DeepEqual(tc.output, output) {
			t.Error("Wrong output")
		}
	}
}

package product_of_array_except_self_238

import (
	"testing"
	"reflect"
)

type testCase struct {
	nums []int
	output []int
}

func Test_productExceptSelf(t *testing.T) {
	testCaseArr := []testCase {
		{
			nums: []int {1,2,3,4},
			output: []int {24,12,8,6},
		},
		{
			nums: []int {-1,1,0,-3,3},
			output: []int {0,0,9,0,0},
		},
	}

	for i, tc := range testCaseArr {
		if !reflect.DeepEqual(productExceptSelf(tc.nums), tc.output) {
			t.Errorf("Wrong output - %d", i)
		}
	}
}
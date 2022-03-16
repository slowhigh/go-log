package maximum_product_subarray_152

import "testing"

type testCase struct {
	nums   []int
	output int
}

func Test_maxProduct(t *testing.T) {
	testCaseArr := []testCase{
		{
			nums:   []int{2, 3, -2, 4},
			output: 6,
		},
		{
			nums:   []int{-2, 0, -1},
			output: 0,
		},
		{
			nums:   []int{-2, 3, -4},
			output: 24,
		},
		{
			nums:   []int{0, 2},
			output: 2,
		},
		{
			nums:   []int{-3, 0, 1, -2},
			output: 1,
		},
		{
			nums:   []int{-1,-2,-3,0},
			output: 6,
		},
	}

	for i, tc := range testCaseArr {
		if maxProduct1(tc.nums) != tc.output {
			t.Errorf("maxProduct1 - Wrong Error - %d", i)
		}

		if maxProduct2(tc.nums) != tc.output {
			t.Errorf("maxProduct2 - Wrong Error - %d", i)
		}
	}
}

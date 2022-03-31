package container_with_most_water_11

import "testing"

type testCase struct {
	height []int
	output int
}

func Test_maxArea(t *testing.T) {
	testCaseArr := []testCase {
		{
			height: []int {1,8,6,2,5,4,8,3,7},
			output: 49,
		},
		{
			height: []int {1,1},
			output: 1,
		},
	}

	for i, tc := range testCaseArr {
		if maxArea(tc.height) != tc.output {
			t.Errorf("Wrong output - %d", i)
		}
	}
}
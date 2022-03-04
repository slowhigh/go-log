package remove_covered_intervals_1288

import "testing"

type testCase struct {
	intervals [][]int
	output    int
}

func Test_removeCoveredIntervals(t *testing.T) {
	testCaseArr := []testCase{
		{
			intervals: [][]int{{1, 4}, {3, 6}, {2, 8}},
			output:    2,
		},
		{
			intervals: [][]int{{1, 4}, {2, 3}},
			output:    1,
		},
	}

	for _, tc := range testCaseArr {
		if removeCoveredIntervals(tc.intervals) != tc.output {
			t.Error("Wrong output")
		}
	}
}

package h_index

import "testing"

type testCase struct {
	citations []int
	result int
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase {
		{
			citations: []int {3, 0, 6, 1, 5},
			result: 3,
		},
		{
			citations: []int {0, 0, 0, 0, 0},
			result: 0,
		},
		{
			citations: []int {6, 6, 6, 6, 6},
			result: 5,
		},
		{
			citations: []int {3},
			result: 1,
		},
		{
			citations: []int {0},
			result: 0,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.citations) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}
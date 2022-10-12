package internal

import "testing"

type testCase struct {
	a      []int
	b      []int
	result int
}

func TestSolution(t *testing.T) {
	testCaseArr := []testCase{
		{
			a:      []int{1, 2, 3, 4},
			b:      []int{-3, -1, 0, 2},
			result: 3,
		},
		{
			a:      []int{-1, 0, 1},
			b:      []int{1, 0, -1},
			result: -2,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.a, tc.b) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

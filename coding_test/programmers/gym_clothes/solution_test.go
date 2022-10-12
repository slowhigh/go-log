package gym_clothes

import "testing"

type testCase struct {
	n       int
	lost    []int
	reserve []int
	result  int
}

func TestSolution(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:       5,
			lost:    []int{2, 4},
			reserve: []int{1, 3, 5},
			result:  5,
		},
		{
			n:       5,
			lost:    []int{2, 4},
			reserve: []int{3},
			result:  4,
		},
		{
			n:       3,
			lost:    []int{3},
			reserve: []int{1},
			result:  2,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.n, tc.lost, tc.reserve) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

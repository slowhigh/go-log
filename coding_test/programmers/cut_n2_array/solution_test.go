package cut_n2_array

import (
	"fmt"
	"testing"
)

type testCase struct {
	n      int
	left   int64
	right  int64
	result []int
}

func TestSolution(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:      3,
			left:   2,
			right:  5,
			result: []int{3, 2, 2, 3},
		},
		{
			n:      4,
			left:   7,
			right:  14,
			result: []int{4, 3, 3, 3, 4, 4, 4, 4},
		},
	}

	for i, tc := range testCaseArr {
		if fmt.Sprint(solution(tc.n, tc.left, tc.right)) != fmt.Sprint(tc.result) {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

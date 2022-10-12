package k_th_number

import (
	"fmt"
	"testing"
)

type testCase struct {
	array    []int
	commands [][]int
	result   []int
}

func TestSolution(t *testing.T) {
	testCaseArr := []testCase{
		{
			array:    []int{1, 5, 2, 6, 3, 7, 4},
			commands: [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}},
			result:   []int{5, 6, 3},
		},
	}

	for i, tc := range testCaseArr {
		if fmt.Sprint(solution(tc.array, tc.commands)) != fmt.Sprint(tc.result) {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

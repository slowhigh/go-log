package matrix_multiplication

import (
	"fmt"
	"testing"
)

type testCase struct {
	arr1   [][]int
	arr2   [][]int
	result [][]int
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		// {
		// 	arr1:   [][]int{{1, 4}, {3, 2}, {4, 1}},
		// 	arr2:   [][]int{{3, 3}, {3, 3}},
		// 	result: [][]int{{15, 15}, {15, 15}, {15, 15}},
		// },
		// {
		// 	arr1:   [][]int{{2, 3, 2}, {4, 2, 4}, {3, 1, 4}},
		// 	arr2:   [][]int{{5, 4, 3}, {2, 4, 1}, {3, 1, 1}},
		// 	result: [][]int{{22, 22, 11}, {36, 28, 18}, {29, 20, 14}},
		// },
		// {
		// 	arr1:   [][]int{{4}},
		// 	arr2:   [][]int{{5}},
		// 	result: [][]int{{20}},
		// },
		{
			arr1:   [][]int{{1,2}},
			arr2:   [][]int{{3}, {4}},
			result: [][]int{{11}},
		},
	}

	for i, tc := range testCaseArr {
		if fmt.Sprint(solution(tc.arr1, tc.arr2)) != fmt.Sprint(tc.result) {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

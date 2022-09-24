package n_least_common_multiple

import "testing"

type testCase struct {
	arr    []int
	result int
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			arr:    []int{2, 6, 8, 14},
			result: 168,
		},
		{
			arr:    []int{1, 2, 3},
			result: 6,
		},
		{
			arr:    []int{3, 7, 11},
			result: 231,
		},
		{
			arr:    []int{7, 7, 7},
			result: 7,
		},
		{
			arr:    []int{2, 4, 8},
			result: 8,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.arr) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

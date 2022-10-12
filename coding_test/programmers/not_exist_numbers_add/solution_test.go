package not_exist_numbers_add

import "testing"

type testCase struct {
	numbers []int
	result  int
}

func TestSolution(t *testing.T) {
	testCaseArr := []testCase{
		{
			numbers: []int{1, 2, 3, 4, 6, 7, 8, 0},
			result:  14,
		},
		{
			numbers: []int{5, 8, 4, 0, 6, 7, 9},
			result:  6,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.numbers) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

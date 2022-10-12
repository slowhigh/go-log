package prime_number_make

import "testing"

type testCase struct {
	nums   []int
	result int
}

func TestSolution(t *testing.T) {
	testCaseArr := []testCase{
		{
			nums:   []int{1, 2, 3, 4},
			result: 1,
		},
		{
			nums:   []int{1, 2, 7, 6, 4},
			result: 4,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.nums) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

package find_the_number_of_remaining_ones

import "testing"

type testCase struct {
	n      int
	result int
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:      10,
			result: 3,
		},
		{
			n:      12,
			result: 11,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.n) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

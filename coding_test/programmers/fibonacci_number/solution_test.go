package fibonacci_number

import "testing"

type testCase struct {
	n      int
	result int
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:      3,
			result: 2,
		},
		{
			n:      5,
			result: 5,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.n) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

package long_jump

import "testing"

type testCase struct {
	n      int
	result int64
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:      4,
			result: 5,
		},
		{
			n:      3,
			result: 3,
		},
		{
			n:      1,
			result: 1,
		},
		{
			n:      2,
			result: 2,
		},
		{
			n:      5,
			result: 8,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.n) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

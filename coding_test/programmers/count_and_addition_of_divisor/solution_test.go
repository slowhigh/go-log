package count_and_addition_of_divisor

import "testing"

type testCase struct {
	left   int
	right  int
	result int
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase {
		{
			left: 1,
			right: 1,
			result: -1,
		},
		{
			left: 13,
			right: 17,
			result: 43,
		},
		{
			left: 24,
			right: 27,
			result: 52,
		},
		{
			left: 5,
			right: 5,
			result: 5,
		},
		{
			left: 25,
			right: 25,
			result: -25,
		},

	}

	for i, tc := range testCaseArr {
		if solution(tc.left, tc.right) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

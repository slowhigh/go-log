package rotate_parentheses

import "testing"

type testCase struct {
	s      string
	result int
}

func TestSolution(t *testing.T) {
	testCaseArr := []testCase{
		{
			s:      "[](){}",
			result: 3,
		},
		{
			s:      "}]()[{",
			result: 2,
		},
		{
			s:      "[)(]",
			result: 0,
		},
		{
			s:      "}}}",
			result: 0,
		},
		{
			s:      "[",
			result: 0,
		},
		{
			s:      ")",
			result: 0,
		},
		{
			s:      "()",
			result: 1,
		},
		{
			s:      "(((",
			result: 0,
		},
		{
			s:      "((()",
			result: 0,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.s) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

package maximum_and_minimum

import "testing"

type testCase struct {
	s string
	r string
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			s: "1 2 3 4",
			r: "1 4",
		},
		{
			s: "-1 -2 -3 -4",
			r: "-4 -1",
		},
		{
			s: "-1 -1",
			r: "-1 -1",
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.s) != tc.r {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

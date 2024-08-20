package contact

import "testing"

type testCase struct {
	s string
	r string
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			s: "10010111",
			r: "NO",
		},
		{
			s: "011000100110001",
			r: "NO",
		},
		{
			s: "0110001011001",
			r: "YES",
		},
	}

	for i, tc := range testCaseArr {
		res := solution(tc.s)

		if res == tc.r {
			t.Logf("PASS - %d(%s)", i, res)
		} else {
			t.Errorf("FAIL - %d(%s)", i, res)
		}
	}
}

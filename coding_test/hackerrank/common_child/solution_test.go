package commonChild

import "testing"

type testCase struct {
	s1     string
	s2     string
	result int32
}

// "ABDF"
// "FBDAN"

func TestCommonChild(t *testing.T) {
	testCaseArr := []testCase {
		{
			s1: "HARRY",
			s2: "SALLY",
			result: 2,
		},
		{
			s1: "AA",
			s2: "BB",
			result: 0,
		},
		{
			s1: "SHINCHAN",
			s2: "NOHARAAA",
			result: 3,
		},
		{
			s1: "1234526412345264",
			s2: "4726R6664726R666",
			result: 6,
		},
		{
			s1: "ABCDEF",
			s2: "FBDAMN",
			result: 2,
		},
	}

	for i, tc := range testCaseArr {
		res := commonChild(tc.s1, tc.s2)

		if res == tc.result {
			t.Logf("PASS - %d(%d)", i, res)
		} else {
			t.Errorf("FAIL - %d(%d)", i, res)
		}
	}
}

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
			s1: "SALLY",
			s2: "HARRY",
			result: 2,
		},
		{
			s1: "AA",
			s2: "BB",
			result: 0,
		},
		{
			s1: "NOHARAAA",
			s2: "SHINCHAN",
			result: 3,
		},
		{
			s1: "4726R6664726R666", // 426666426666
			s2: "1234526412345264", // 2426424264
			result: 6,
		},
		{
			s1: "4726R666", // 426666
			s2: "12345264", // 24264
			result: 3,
		},
		{
			s1: "FBDAMN",
			s2: "ABCDEF",
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

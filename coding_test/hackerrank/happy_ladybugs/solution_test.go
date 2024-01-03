package happy_ladybugs

import (
	"testing"
)

type testCase struct {
	b      string
	result string
}

func Test_happyLadybugs(t *testing.T) {
	testCaseArr := []testCase{
		{
			b:      "AABCBC",
			result: "NO",
		},
		{
			b:      "RBY_YBR",
			result: "YES",
		},
		{
			b:      "X_Y__X",
			result: "NO",
		},
		{
			b:      "__",
			result: "YES",
		},
		{
			b:      "B_RRBR",
			result: "YES",
		},
	}

	for i, tc := range testCaseArr {
		res := happyLadybugs(tc.b)

		if res == tc.result {
			t.Logf("PASS - %d (%s)", i, res)
		} else {
			t.Errorf("FAIL - %d (%s)", i, res)
		}
	}
}

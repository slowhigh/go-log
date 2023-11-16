package counting_valleys

import "testing"

type testCase struct {
	steps  int32
	path   string
	result int32
}

func Test_countingValleys(t *testing.T) {
	testCaseArr := []testCase{
		{
			steps:  8,
			path:   "UDDDUDUU",
			result: 1,
		},
		{
			steps:  12,
			path:   "DDUUDDUDUUUD",
			result: 2,
		},
	}

	for i, tc := range testCaseArr {
		res := countingValleys(tc.steps, tc.path)
		if res == tc.result {
			t.Logf("PASS - %d (%d)", i, res)
		} else {
			t.Errorf("FAIL - %d (%d)", i, res)
		}
	}
}

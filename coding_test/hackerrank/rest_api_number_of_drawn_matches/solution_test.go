package rest_api_number_of_drawn_matches

import "testing"

type testCase struct {
	year   int
	result int
}

func Test_getNumDraws(t *testing.T) {
	testCaseArr := []testCase{
		{
			year:   2011,
			result: 516,
		},
	}

	for i, tc := range testCaseArr {
		output := getNumDraws(tc.year)

		if output == tc.result {
			t.Logf("PASS - %d / Expected Output=%d / Your Output=%d)", i, tc.result, output)
		} else {
			t.Errorf("FAIL - %d / Expected Output=%d / Your Output=%d)", i, tc.result, output)
		}
	}
}

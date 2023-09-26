package rest_api_football_competition_winners_goals

import "testing"

type testCase struct {
	competition string
	year        int
	result      int
}

func Test_getWinnerTotalGoals(t *testing.T) {
	testCaseArr := []testCase{
		{
			competition: "UEFA Champions League",
			year:        2011,
			result:      28,
		},
	}

	for i, tc := range testCaseArr {
		result := getWinnerTotalGoals(tc.competition, tc.year)

		if result == tc.result {
			t.Logf("PASS - %d / Expected Output=%d / Your Output=%d)", i, tc.result, result)
		} else {
			t.Errorf("FAIL - %d / Expected Output=%d / Your Output=%d)", i, tc.result, result)
		}
	}
}

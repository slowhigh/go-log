package rest_api_total_goals_by_a_team

import "testing"

type testCase struct {
	team  string
	year  int32
	goals int32
}

func Test_GetTotalGoals(t *testing.T) {
	testCaseArr := []testCase{
		{
			team:  "Barcelona",
			year:  2011,
			goals: 35,
		},
	}

	for i, tc := range testCaseArr {
		goals := GetTotalGoals(tc.team, tc.year)
		if goals == tc.goals {
			t.Logf("PASS - %d ( %d vs %d )", i, goals, tc.goals)
		} else {
			t.Errorf("FAIL - %d ( %d vs %d )", i, goals, tc.goals)
		}
	}
}

func Test_GetTotalGoals2(t *testing.T) {
	testCaseArr := []testCase{
		{
			team:  "Barcelona",
			year:  2011,
			goals: 35,
		},
	}

	for i, tc := range testCaseArr {
		goals := GetTotalGoals2(tc.team, tc.year)
		if goals == tc.goals {
			t.Logf("PASS - %d ( %d vs %d )", i, goals, tc.goals)
		} else {
			t.Errorf("FAIL - %d ( %d vs %d )", i, goals, tc.goals)
		}
	}
}


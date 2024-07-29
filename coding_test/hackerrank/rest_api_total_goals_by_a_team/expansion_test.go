package restApiTotalGoalsByATeam

import "testing"

func Test_GetAllTeamTotalGoals(t *testing.T) {
	teams, err := GetAllTeamTotalGoals()
	if err != nil {
		t.Error(err)
	}

	if len(*teams) == 187 {
		t.Logf("PASS - len(teams) = %d", len(*teams))
	} else {
		t.Errorf("FAIL - len(teams) = %d", len(*teams))
	}
}

package theGridSearch

import "testing"

type testCase struct {
	g      []string
	p      []string
	result string
}

func Test_gridSearch(t *testing.T) {
	testCaseArr := []testCase{
		{
			g: []string{
				"7283455864",
				"6731158619",
				"8988242643",
				"3830589324",
				"2229505813",
				"5633845374",
				"6473530293",
				"7053106601",
				"0834282956",
				"4607924137",
			},
			p: []string{
				"9505",
				"3845",
				"3530",
			},
			result: "YES",
		},
	}

	for i, tc := range testCaseArr {
		if gridSearch(tc.g, tc.p) == tc.result {
			t.Logf("PASS - %d", i)
		} else {
			t.Errorf("FAIL - %d", i)
		}
	}
}

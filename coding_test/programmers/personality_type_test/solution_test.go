package personality_type_test

import "testing"

type testCase struct {
	survey  []string
	choices []int
	result  string
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			[]string{"AN", "CF", "MJ", "RT", "NA"},
			[]int{5, 3, 2, 7, 5},
			"TCMA",
		},
		{
			[]string{"TR", "RT", "TR"},
			[]int{7, 1, 3},
			"RCJA",
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.survey, tc.choices) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

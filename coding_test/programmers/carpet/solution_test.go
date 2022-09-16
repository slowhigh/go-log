package carpet

import (
	"fmt"
	"testing"
)

type testCase struct {
	brown  int
	yellow int
	result []int
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			brown:  10,
			yellow: 2,
			result: []int{4, 3},
		},
		{
			brown:  8,
			yellow: 1,
			result: []int{3, 3},
		},
		{
			brown:  24,
			yellow: 24,
			result: []int{8, 6},
		},
	}

	for i, tc := range testCaseArr {
		if fmt.Sprint(solution(tc.brown, tc.yellow)) != fmt.Sprint(tc.result) {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

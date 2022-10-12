package negative_positive_number_add

import "testing"

type testCase struct {
	absolutes []int
	signs     []bool
	result    int
}

func TestSolution(t *testing.T) {
	testCaseArr := []testCase{
		{
			absolutes: []int{4, 7, 12},
			signs:     []bool{true, false, true},
			result:    9,
		},
		{
			absolutes: []int{1, 2, 3},
			signs:     []bool{false, false, true},
			result:    0,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.absolutes, tc.signs) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

package lotto_highest_and_lowest_ranking

import (
	"fmt"
	"testing"
)

type testCase struct {
	lottos   []int
	win_nums []int
	result   []int
}

func TestSolution(t *testing.T) {
	testCaseArr := []testCase{
		{
			lottos:   []int{44, 1, 0, 0, 31, 25},
			win_nums: []int{31, 10, 45, 1, 6, 19},
			result:   []int{3, 5},
		},
		{
			lottos:   []int{0, 0, 0, 0, 0, 0},
			win_nums: []int{38, 19, 20, 40, 15, 25},
			result:   []int{1, 6},
		},
		{
			lottos:   []int{45, 4, 35, 20, 3, 9},
			win_nums: []int{20, 9, 3, 45, 4, 35},
			result:   []int{1, 1},
		},
	}

	for i, tc := range testCaseArr {
		if fmt.Sprint(solution(tc.lottos, tc.win_nums)) != fmt.Sprint(tc.result) {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

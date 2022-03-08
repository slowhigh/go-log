package best_time_to_buy_and_sell_stock_121

import (
	"testing"
)

type testCase struct {
	prices []int
	output int
}

func Test_maxProfit(t *testing.T) {
	testCaseArr := []testCase {
		{
			prices: []int {7,2,4,1},
			output: 2,
		},
		{
			prices: []int {1,1,0},
			output: 0,
		},
		{
			prices: []int {7,1,5,3,6,4},
			output: 5,
		},
		{
			prices: []int {7,6,4,3,1},
			output: 0,
		},
		{
			prices: []int {2,4,1},
			output: 2,
		},
		{
			prices: []int {4,2,1},
			output: 0,
		},
	}

	for ti, tc := range testCaseArr {
		if maxProfit(tc.prices) != tc.output {
			t.Errorf("Wrong output - %d", ti)
		}
	}
}
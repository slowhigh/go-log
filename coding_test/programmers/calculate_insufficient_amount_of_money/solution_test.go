package calculate_insufficient_amount_of_money

import "testing"

type testCase struct {
	price  int
	money  int
	count  int
	result int64
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			price:  3,
			money:  20,
			count:  4,
			result: 10,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.price, tc.money, tc.count) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}

	for i, tc := range testCaseArr {
		if solution2(tc.price, tc.money, tc.count) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

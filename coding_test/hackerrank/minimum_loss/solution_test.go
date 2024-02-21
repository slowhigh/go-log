package minimum_loss

import "testing"

type testCase struct {
	price  []int64
	result int32
}

func Test_minimumLoss(t *testing.T) {
	testCaseArr := []testCase{
		{
			price:  []int64{5, 10, 3},
			result: 2,
		},
		{
			price:  []int64{20, 7, 8, 2, 5},
			result: 2,
		},
	}

	for i, tc := range testCaseArr {
		res := minimumLoss(tc.price)

		if res == tc.result {
			t.Logf("PASS - %d(%d)", i, res)
		} else {
			t.Errorf("FAIL - %d(%d)", i, res)
		}
	}
}

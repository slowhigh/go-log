package pickingNumbers

import "testing"

// 2 <= len(a) <= 100
// 0 < a[i] < 100
// 2 <= result
type testCase struct {
	a      []int32
	result int32
}

func Test_pickingNumbers(t *testing.T) {
	testCaseArr := []testCase{
		{
			a:      []int32{1, 3},
			result: 1,
		},
		{
			a:      []int32{1, 1, 2, 2, 4, 4, 5, 5, 5},
			result: 5,
		},
		{
			a:      []int32{4, 6, 5, 3, 3, 1},
			result: 3,
		},
		{
			a:      []int32{1, 2, 2, 3, 1, 2},
			result: 5,
		},
	}

	for i, tc := range testCaseArr {
		res := pickingNumbers(tc.a)
		if res == tc.result {
			t.Logf("PASS - %d (%d)", i, res)
		} else {
			t.Errorf("FAIL - %d (%d)", i, res)
		}
	}
}

package strange_counter

import "testing"

type testCase struct {
	t      int64
	result int64
}

func Test_strangeCounter(t *testing.T) {
	testCaseArr := []testCase{
		{
			t:      11,
			result: 6,
		},
		{
			t:      4,
			result: 6,
		},
	}

	for i, tc := range testCaseArr {
		res := strangeCounter(tc.t)

		if res == tc.result {
			t.Logf("PASS - %d (%d)", i, res)
		} else {
			t.Errorf("FAIL - %d (%d)", i, res)
		}

	}
}

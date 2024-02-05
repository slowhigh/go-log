package equal

import (
	"testing"
)

type testCase struct {
	arr    []int32
	result int32
}

func Test_equal(t *testing.T) {
	testCaseArr := []testCase{
		{
			arr:    []int32{2, 2, 3, 7},
			result: 2,
		},
		{
			arr:    []int32{10, 7, 12},
			result: 3,
		},
	}

	for i, tc := range testCaseArr {
		res := equal(tc.arr)

		if res == tc.result {
			t.Logf("PASS - %d(%d)", i, res)
		} else {
			t.Errorf("FAIL - %d(%d)", i, res)
		}
	}
}

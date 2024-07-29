package larrySArray

import (
	"testing"
)

type testCase struct {
	A      []int32
	result string
}

func Test_larrysArray(t *testing.T) {
	arr := make([]int32, 1000)
	for i := range arr {
		arr[i] = int32(i + 1)
	}

	testCaseArr := []testCase{
		{
			A:      arr,
			result: "YES",
		},
		{
			A:      []int32{3, 1, 2},
			result: "YES",
		},
		{
			A:      []int32{1, 3, 4, 2},
			result: "YES",
		},
		{
			A:      []int32{1, 2, 3, 5, 4},
			result: "NO",
		},
		{
			A:      []int32{1, 6, 5, 2, 3, 4},
			result: "NO",
		},
		{
			A:      []int32{3, 1, 2, 4},
			result: "YES",
		},
	}

	for i, tc := range testCaseArr {
		res := larrysArray(tc.A)

		if tc.result == res {
			t.Logf("PASS - %d(%s)", i, res)
		} else {
			t.Errorf("FAIL - %d(%s)", i, res)
		}
	}
}

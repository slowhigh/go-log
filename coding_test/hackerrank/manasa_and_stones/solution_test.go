package manasa_and_stones

import (
	"fmt"
	"testing"
)

// 2       T = 2 (test cases)
// 3       n = 3 (test case 1)
// 1       a = 1
// 2       b = 2
// 4       n = 4 (test case 2)
// 10      a = 10
// 100     b = 100

// 2 3 4
// 30 120 210 300

type testCase struct {
	n      int32
	a      int32
	b      int32
	result []int32
}

func Test_stones(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:      3,
			a:      1,
			b:      2,
			result: []int32{2, 3, 4},
		},
		{
			n:      4,
			a:      10,
			b:      100,
			result: []int32{30, 120, 210, 300},
		},
	}

	for i, tc := range testCaseArr {
		res := fmt.Sprintf("%+v", stones(tc.n, tc.a, tc.b))

		if res == fmt.Sprintf("%+v", tc.result) {
			t.Logf("PASS - %d (%s)", i, res)
		} else {
			t.Errorf("FAIL - %d (%s)", i, res)
		}
	}
}

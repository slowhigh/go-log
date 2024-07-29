package closestNumbers

import (
	"fmt"
	"testing"
)

// -20 -3916237 -357920 -3620601 7374819 -7330761 30 6246457 -6461594 266854
// -20 30

// -20 -3916237 -357920 -3620601 7374819 -7330761 30 6246457 -6461594 266854 -520 -470
// -520 -470 -20 30

// 5 4 3 2
// 2 3 3 4 4 5

type testCase struct {
	arr    []int32
	result []int32
}

func Test_closestNumbers(t *testing.T) {
	testCaseArr := []testCase{
		{
			arr:    []int32{-20, -3916237, -357920, -3620601, 7374819, -7330761, 30, 6246457, -6461594, 266854},
			result: []int32{-20, 30},
		},
		{
			arr:    []int32{-20, -3916237, -357920, -3620601, 7374819, -7330761, 30, 6246457, -6461594, 266854, -520, -470},
			result: []int32{-520, -470, -20, 30},
		},
		{
			arr:    []int32{5, 4, 3, 2},
			result: []int32{2, 3, 3, 4, 4, 5},
		},
	}

	for i, tc := range testCaseArr {
		res := fmt.Sprintf("%+v", closestNumbers(tc.arr))

		if res == fmt.Sprintf("%+v", tc.result) {
			t.Logf("PASS - %d (%s)", i, res)
		} else {
			t.Errorf("Fail - %d (%s)", i, res)
		}
	}
}

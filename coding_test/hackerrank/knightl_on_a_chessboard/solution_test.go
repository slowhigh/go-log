package knightlOnAChessboard

import (
	"fmt"
	"testing"
)

type testCase struct {
	n       int32
	results [][]int32
}

func Test_knightlOnAChessboard(t *testing.T) {
	testCaseArr := []testCase{
		{
			n: 5,
			results: [][]int32{
				{4, 4, 2, 8},
				{4, 2, 4, 4},
				{2, 4, -1, -1},
				{8, 4, -1, 1},
			},
		},
	}

	for i, tc := range testCaseArr {
		res := fmt.Sprintf("%+v", knightlOnAChessboard(tc.n))

		if res == fmt.Sprintf("%+v", tc.results) {
			t.Logf("PASS - %d(%s)", i, res)
		} else {
			t.Errorf("FAIL - %d(%s)", i, res)
		}
	}
}

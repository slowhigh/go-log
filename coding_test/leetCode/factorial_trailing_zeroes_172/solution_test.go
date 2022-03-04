package factorial_trailing_zeroes_172

import (
	"testing"
)

type testCase struct {
	n      int
	output int
}

func Test_trailingZeroes(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:      3,
			output: 0,
		},
		{
			n:      5,
			output: 1,
		},
		{
			n:      0,
			output: 0,
		},
	}

	for _, tc := range testCaseArr {
		if tc.output != trailingZeroes(tc.n) {
			t.Error("Wrong output")
		}
	}
}

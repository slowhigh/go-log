package roman_to_integer_13

import (
	"testing"
)

type testCase struct {
	input  string
	output int
}

func Test_romanToInt(t *testing.T) {
	testCaseArr := []testCase{
		{
			input:  "III",
			output: 3,
		},
		{
			input:  "LVIII",
			output: 58,
		},
		{
			input:  "MCMXCIV",
			output: 1994,
		},
	}

	for _, tc := range testCaseArr {
		if romanToInt(tc.input) != tc.output {
			t.Error("Wrong output")
		}
	}
}

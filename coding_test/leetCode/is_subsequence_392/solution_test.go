package is_subsequence_392

import (
	"testing"
)

type testCase struct {
	input_1 string
	input_2 string
	output  bool
}

func TestIsSubsequence(t *testing.T) {
	testCaseArr := []testCase{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
	}

	for _, testCase := range testCaseArr {
		if testCase.output != isSubsequence(testCase.input_1, testCase.input_2) {
			t.Error("Wrong result")
		}
	}
}

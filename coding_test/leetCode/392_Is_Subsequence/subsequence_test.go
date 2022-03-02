package subsequence_test

import (
	"subsequence"
	"testing"
)

// Example 1:
// Input: s = "abc", t = "ahbgdc"
// Output: true

// Example 2:
// Input: s = "axc", t = "ahbgdc"
// Output: false

type testCase struct {
	input_1 string
	input_2 string
	output bool
}

func TestIsSubsequence(t *testing.T) {
	testCaseArr := []testCase {
		testCase{ "abc", "ahbgdc", true },
		testCase{ "axc", "ahbgdc", false },
	}

	for _, testCase := range testCaseArr {
		if testCase.output != subsequence.IsSubsequence(testCase.input_1, testCase.input_2) {
			t.Error("Wrong result")
		}
	}
}

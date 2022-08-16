package first_unique_character_in_a_string_387

import "testing"

type testCase struct {
	input  string
	output int
}

func Test_firstUniqChar(t *testing.T) {
	testCaseArr := []testCase{
		{
			input: "leetcode",
			output: 0,
		},
		{
			input: "loveleetcode",
			output: 2,
		},
		{
			input: "aabb",
			output: -1,
		},
	}

	for i, tc := range testCaseArr {
		if firstUniqChar(tc.input) != tc.output {
			t.Errorf("Wrong output - %d", i)
		}
	}
}

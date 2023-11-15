package sherlock_and_anagrams

import "testing"

type testCase struct {
	s      string
	result int32
}

func Test_sherlockAndAnagrams(t *testing.T) {
	testCaseArr := []testCase{
		{
			s: "abba",
			result: 4,
		},
		{
			s: "abcd",
			result: 0,
		},
		{
			s: "ifailuhkqq",
			result: 3,
		},
		{
			s: "kkkk",
			result: 10,
		},
		{
			s: "cdcd",
			result: 5,
		},
	}

	for i, tc := range testCaseArr {
		if
	}
}

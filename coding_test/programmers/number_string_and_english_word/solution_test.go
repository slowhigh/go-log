package number_string_and_english_word

import "testing"

type testCase struct {
	s      string
	result int
}

func TestSolution(t *testing.T) {
	testCaseArr := []testCase{
		{
			s:      "one4seveneight",
			result: 1478,
		},
		{
			s:      "23four5six7",
			result: 234567,
		},
		{
			s:      "2three45sixseven",
			result: 234567,
		},
		{
			s:      "123",
			result: 123,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.s) != tc.result {
			t.Errorf("Wrong result 1-%d", i)
		}
	}
	for i, tc := range testCaseArr {
		if solution2(tc.s) != tc.result {
			t.Errorf("Wrong result 2-%d", i)
		}
	}
}

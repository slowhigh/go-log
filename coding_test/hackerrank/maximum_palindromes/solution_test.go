package maximum_palindromes

import "testing"

type testCase struct {
	s      string
	l      int32
	r      int32
	result int32
}

func Test_answerQuery(t *testing.T) {
	testCaseArr := []testCase{
		{
			s: "aaaaabc",
			l: 1,
			r: 7,
			result: 3,
		},
		// {
		// 	s: "madamimadam",
		// 	l: 4,
		// 	r: 7,
		// 	result: 2,
		// },
		// {
		// 	s: "week",
		// 	l: 1,
		// 	r: 4,
		// 	result: 2,
		// },
		// {
		// 	s: "week",
		// 	l: 2,
		// 	r: 3,
		// 	result: 1,
		// },
		// {
		// 	s: "abcde",
		// 	l: 1,
		// 	r: 5,
		// 	result: 0,
		// },
		// {
		// 	s: "aaaaa",
		// 	l: 1,
		// 	r: 5,
		// 	result: 1,
		// },
	}

	for i, tc := range testCaseArr {
		initialize(tc.s)
		res := answerQuery(tc.l, tc.r)

		if res == tc.result {
			t.Logf("PASS - %d (%d)", i, res)
		} else {
			t.Errorf("FAIL - %d (%d)", i, res)
		}
	}
}

package highestValuePalindrome

import "testing"

type testCase struct {
	s      string
	n      int32
	k      int32
	result string
}

func Test_highestValuePalindrome(t *testing.T) {
	testCaseArr := []testCase{
		{
			s:      "1231",
			n:      4,
			k:      3,
			result: "9339",
		},
		{
			s:      "12321",
			n:      5,
			k:      1,
			result: "12921",
		},
		{
			s:      "3943",
			n:      4,
			k:      1,
			result: "3993",
		},
		{
			s:      "092282",
			n:      6,
			k:      3,
			result: "992299",
		},
		{
			s:      "0011",
			n:      4,
			k:      1,
			result: "-1",
		},
	}

	for i, tc := range testCaseArr {
		res := highestValuePalindrome(tc.s, tc.n, tc.k)

		if res == tc.result {
			t.Logf("PASS - %d (%s)", i, res)
		} else {
			t.Errorf("FAIL - %d (%s)", i, res)
		}
	}
}

package make_jadencase_string

import (
	"testing"
)

type testCase struct {
	s string
	r string
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			s: "3people  unFollowed     me",
			r: "3people  Unfollowed     Me",
		},
		{
			s: "for the last week",
			r: "For The Last Week",
		},
	}

	for i, tc := range testCaseArr {
		s := solution(tc.s)

		if s != tc.r {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

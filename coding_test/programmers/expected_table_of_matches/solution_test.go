package expected_table_of_matches

import "testing"

type testCase struct {
	N      int
	A      int
	B      int
	answer int
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			N:      64,
			A:      1,
			B:      64,
			answer: 6,
		},
		{
			N:      8,
			A:      4,
			B:      7,
			answer: 3,
		},
		{
			N:      8,
			A:      1,
			B:      2,
			answer: 1,
		},
		{
			N:      8,
			A:      1,
			B:      3,
			answer: 2,
		},
		{
			N:      8,
			A:      1,
			B:      8,
			answer: 3,
		},
		{
			N:      8,
			A:      1,
			B:      7,
			answer: 3,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.N, tc.A, tc.B) != tc.answer {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

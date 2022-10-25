package find_the_prime_number_in_k_series

import "testing"

type testCase struct {
	n      int
	k      int
	result int
}

func TestSolution(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:      437674,
			k:      3,
			result: 3,
		},
		{
			n:      110011,
			k:      10,
			result: 2,
		},
		{
			n:      2,
			k:      3,
			result: 1,
		},
		{
			n:      3,
			k:      3,
			result: 0,
		},
		{
			n:      1,
			k:      3,
			result: 0,
		},
		{
			n:      10,
			k:      10,
			result: 0,
		},
		{
			n:      33,
			k:      3,
			result: 1,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.n, tc.k) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

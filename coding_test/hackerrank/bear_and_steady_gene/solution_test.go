package bearAndSteadyGene

import "testing"

// GAAATAAA
// 5
// TGATGCCGTCCCCTCAACTTGAGTGCTCCTAATGCGTTGC
// 5

type testCase struct {
	gene   string
	result int32
}

func Test_steadyGene(t *testing.T) {
	testCaseArr := []testCase{
		// {
		// 	gene:   "GAAATAAA",
		// 	result: 5,
		// },
		{
			gene:   "TGATGCCGTCCCCTCAACTTGAGTGCTCCTAATGCGTTGC",
			result: 5,
		},
	}

	for i, tc := range testCaseArr {
		res := steadyGene(tc.gene)

		if res == tc.result {
			t.Logf("PASS - %d(%d)", i, res)
		} else {
			t.Errorf("FAIL - %d(%d)", i, res)
		}
	}
}

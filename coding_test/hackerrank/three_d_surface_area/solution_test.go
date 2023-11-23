package three_d_surface_area

import "testing"

type testCase struct {
	a      [][]int32
	result int32
}

func Test_surfaceArea(t *testing.T) {
	testCaseArr := []testCase{
		{
			a: [][]int32{
				{1, 3, 4},
				{2, 2, 3},
				{1, 2, 4},
			},
			result: 60,
		},
		{
			a: [][]int32{
				{1},
			},
			result: 6,
		},
	}

	for i, tc := range testCaseArr {
		res := surfaceArea(tc.a)
		if res == tc.result {
			t.Logf("PASS - %d (%d)", i, res)
		} else {
			t.Errorf("FAIL - %d (%d)", i, res)
		}
	}
}

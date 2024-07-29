package kruskalMstReallySpecialSubtree

import "testing"

type testCase struct {
	gNodes  int32
	gFrom   []int32
	gTo     []int32
	gWeight []int32
	result  int32
}

func Test_kruskals(t *testing.T) {
	testCaseArr := []testCase{
		{
			gNodes:  3,
			gFrom:   []int32{1, 1, 2},
			gTo:     []int32{2, 2, 1},
			gWeight: []int32{2, 3, 4},
			result:  2,
		},
		{
			gNodes:  4,
			gFrom:   []int32{1, 1, 4, 2, 3, 3},
			gTo:     []int32{2, 3, 1, 4, 2, 4},
			gWeight: []int32{5, 3, 6, 7, 4, 5},
			result:  12,
		},
		{
			gNodes:  5,
			gFrom:   []int32{1, 1, 3, 1, 1, 2, 3, 4},
			gTo:     []int32{2, 3, 1, 4, 5, 3, 4, 5},
			gWeight: []int32{20, 50, 60, 70, 90, 30, 40, 60},
			result:  150,
		},

		{
			gNodes:  4,
			gFrom:   []int32{1, 3, 4, 1, 3},
			gTo:     []int32{2, 2, 3, 4, 1},
			gWeight: []int32{1, 150, 99, 100, 200},
			result:  200,
		},
	}

	for i, tc := range testCaseArr {
		if kruskals(tc.gNodes, tc.gFrom, tc.gTo, tc.gWeight) == tc.result {
			t.Logf("PASS - %d", i)
		} else {
			t.Errorf("FAIL - %d", i)
		}
	}
}

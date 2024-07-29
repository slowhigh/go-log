package specialSubtree

import "testing"

type testCase struct {
	n      int32
	edges  [][]int32
	start  int32
	result int32
}

func Test_prims(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:      5,
			edges:  [][]int32{{1, 2, 3}, {1, 3, 4}, {4, 2, 6}, {5, 2, 2}, {2, 3, 5}, {3, 5, 7}},
			start:  1,
			result: 15,
		},
	}

	for i, tc := range testCaseArr {
		result := prims(tc.n, tc.edges, tc.start)

		if result == tc.result {
			t.Logf("PASS - %d", i)
		} else {
			t.Errorf("FAIL - %d (%d vs %d)", i, result, tc.result)
		}
	}
}

package even_tree

import "testing"

type testCase struct {
	t_nodes int32
	t_edges int32
	t_from  []int32
	t_to    []int32
	result  int32
}

func Test_evenForest(t *testing.T) {
	testCaseArr := []testCase{
		{
			t_nodes: 4,
			t_edges: 3,
			t_from:  []int32{1, 1, 3},
			t_to:    []int32{2, 3, 4},
			result:  1,
		},
		{
			t_nodes: 10,
			t_edges: 9,
			t_from:  []int32{2, 3, 4, 5, 6, 7, 8, 9, 10},
			t_to:    []int32{1, 1, 3, 2, 1, 2, 6, 8, 8},
			result:  2,
		},
		{
			t_nodes: 10,
			t_edges: 9,
			t_from:  []int32{1, 2, 3, 4, 5, 6, 7, 8, 9},
			t_to:    []int32{2, 3, 4, 5, 6, 7, 8, 9, 10},
			result:  4,
		},
		{
			t_nodes: 30,
			t_edges: 29,
			t_from:  []int32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30,},
			t_to:    []int32{1, 2, 3, 2, 4, 4, 1, 5, 4, 4, 8, 2, 2, 8, 10, 1, 17, 18, 4, 15, 20, 2, 12, 21, 17, 5, 27, 4, 25,},
			result:  11,
		},
	}

	for i, tc := range testCaseArr {
		if evenForest(tc.t_nodes, tc.t_edges, tc.t_from, tc.t_to) == tc.result {
			t.Logf("PASS - %d", i)
		} else {
			t.Errorf("FAIL - %d", i)
		}
	}
}

package breadth_first_search_shortest_reach

import (
	"fmt"
	"testing"
)

type testCase struct {
	n      int32
	m      int32
	edges  [][]int32
	s      int32
	result []int32
}

func Test_bfs(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:      5,
			m:      3,
			edges:  [][]int32{{1, 2}, {1, 3}, {3, 4}},
			s:      1,
			result: []int32{6, 6, 12, -1},
		},
		{
			n:      4,
			m:      2,
			edges:  [][]int32{{1, 2}, {1, 3}},
			s:      1,
			result: []int32{6, 6, -1},
		},
		{
			n:      3,
			m:      1,
			edges:  [][]int32{{2, 3}},
			s:      2,
			result: []int32{-1, 6},
		},
	}

	for i, tc := range testCaseArr {
		if fmt.Sprint(bfs(tc.n, tc.m, tc.edges, tc.s)) == fmt.Sprint(tc.result) {
			t.Logf("PASS - %d", i)
		} else {
			t.Errorf("FAIL - %d", i)
		}
	}
}

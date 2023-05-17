package roads_and_libraries

import "testing"

type testCase struct {
	n      int32
	c_lib  int32
	c_road int32
	cities [][]int32
	result int64
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:      3,
			c_lib:  2,
			c_road: 1,
			cities: [][]int32{{1, 2}, {3, 1}, {2, 3}},
			result: 4,
		},
		{
			n:      6,
			c_lib:  2,
			c_road: 5,
			cities: [][]int32{{1, 3}, {3, 4}, {2, 4}, {1, 2}, {2, 3}, {5, 6}},
			result: 12,
		},
		{
			n:      6,
			c_lib:  2,
			c_road: 3,
			cities: [][]int32{{1, 2}, {1, 3}, {4, 5}, {4, 6}},
			result: 12,
		},
		{
			n:      5,
			c_lib:  6,
			c_road: 1,
			cities: [][]int32{{1, 2}, {1, 3}, {1, 4}},
			result: 15,
		},
	}

	for i, tc := range testCaseArr {
		if roadsAndLibraries(tc.n, tc.c_lib, tc.c_road, tc.cities) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

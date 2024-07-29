package snakesAndLaddersTheQuickestWayUp

import "testing"

type testCase struct {
	ladders [][]int32
	snakes  [][]int32
	result  int32
}

func Test_quickestWayUp(t *testing.T) {
	testCaseArr := []testCase{
		{
			ladders: [][]int32{{3, 90}},
			snakes:  [][]int32{{99, 10}, {97, 20}, {98, 30}, {96, 40}, {95, 50}, {94, 60}, {93, 70}},
			result:  -1,
		},
		{
			ladders: [][]int32{{3, 54}, {37, 100}},
			snakes:  [][]int32{{56, 33}},
			result:  3,
		},
		{
			ladders: [][]int32{{3, 57}, {8, 100}},
			snakes:  [][]int32{{88, 44}},
			result:  2,
		},
		{
			ladders: [][]int32{{7, 98}},
			snakes:  [][]int32{{99, 1}},
			result:  2,
		},
		{
			ladders: [][]int32{{32, 62}, {42, 68}, {12, 98}},
			snakes:  [][]int32{{95, 13}, {97, 25}, {93, 37}, {79, 27}, {75, 19}, {49, 47}, {67, 17}},
			result:  3,
		},
		{
			ladders: [][]int32{{8, 52}, {6, 80}, {26, 42}, {2, 72}},
			snakes:  [][]int32{{51, 19}, {39, 11}, {37, 29}, {81, 3}, {59, 5}, {79, 23}, {53, 7}, {43, 33}, {77, 21}},
			result:  5,
		},
	}

	for i, tc := range testCaseArr {
		if quickestWayUp(tc.ladders, tc.snakes) == tc.result {
			t.Logf("PASS - %d", i)
		} else {
			t.Errorf("FAIL - %d", i)
		}
	}
}

package synchronousShopping

import "testing"

type testCase struct {
	n       int32
	k       int32
	centers []string
	roads   [][]int32
	result  int32
}

func Test_shop(t *testing.T) {
	testCaseArr := []testCase{
		{
			n: 6,
			k: 4,
			centers: []string{
				"1 2",
				"1 2",
				"1 1",
				"2 3 4",
				"2 3 4",
				"1 4",
			},
			roads: [][]int32{
				{5, 4, 646},
				{4, 1, 997},
				{2, 1, 881},
				{2, 6, 114},
				{3, 1, 46},
			},
			result: 2989,
		},
		{
			n: 5,
			k: 3,
			centers: []string{
				"1 1",
				"2 1 2",
				"2 2 3",
				"1 2",
				"1 2"},
			roads: [][]int32{
				{1, 2, 10},
				{1, 3, 15},
				{1, 4, 1},
				{2, 3, 10},
				{3, 5, 5}},
			result: 20,
		},
		{
			n: 5,
			k: 5,
			centers: []string{
				"1 1",
				"1 2",
				"1 3",
				"1 4",
				"1 5"},
			roads: [][]int32{
				{1, 2, 10},
				{1, 3, 10},
				{2, 4, 10},
				{3, 5, 10},
				{4, 5, 10}},
			result: 30,
		},
		{
			n: 6,
			k: 3,
			centers: []string{
				"2 1 2",
				"1 3",
				"0",
				"2 1 3",
				"1 2",
				"1 3"},
			roads: [][]int32{
				{1, 2, 572},
				{4, 2, 913},
				{2, 6, 220},
				{1, 3, 579},
				{2, 3, 808},
				{5, 3, 298},
				{6, 1, 927},
				{4, 5, 171},
				{1, 5, 671},
				{2, 5, 463},
			},
			result: 792,
		},
	}

	for i, tc := range testCaseArr {
		if shop(tc.n, tc.k, tc.centers, tc.roads) == tc.result {
			t.Logf("PASS - %d", i)
		} else {
			t.Errorf("FAIL - %d", i)
		}
	}
}

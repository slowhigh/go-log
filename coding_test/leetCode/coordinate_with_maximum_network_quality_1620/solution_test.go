package coordinate_with_maximum_network_quality_1620

import (
	"reflect"
	"testing"
)

type testCase struct {
	tower  [][]int
	radius int
	output []int
}

func TestBestCoordinate(t *testing.T) {
	testCaseArr := []testCase{
		{
			tower:  [][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}},
			radius: 2,
			output: []int{2, 1},
		},
		{
			tower:  [][]int{{23, 11, 21}},
			radius: 9,
			output: []int{23, 11},
		},
		{
			tower:  [][]int{{1, 2, 13}, {2, 1, 7}, {0, 1, 9}},
			radius: 2,
			output: []int{1, 2},
		},
	}

	for _, tc := range testCaseArr {
		result := bestCoordinate(tc.tower, tc.radius)

		if !reflect.DeepEqual(tc.output, result) {
			t.Error("Wrong output")
		}
	}
}

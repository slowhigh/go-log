package journey_to_the_moon

import "testing"

type testCase struct {
	n         int32
	astronaut [][]int32
	result    int64
}

func Test_journeyToMoon(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:         5,
			astronaut: [][]int32{{0, 1}, {2, 3}, {0, 4}},
			result:    6,
		},
		{
			n:         4,
			astronaut: [][]int32{{0, 2}},
			result:    5,
		},
	}

	for i, tc := range testCaseArr {
		if journeyToMoon(tc.n, tc.astronaut) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

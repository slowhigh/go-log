package camouflage

import "testing"

type testCase struct {
	clothes [][]string
	result  int
}

func TestSolution(t *testing.T) {
	testCaseArr := []testCase{
		{
			clothes: [][]string{{"yellow_hat", "headgear"}, {"blue_sunglasses", "eyewear"}, {"green_turban", "headgear"}},
			result:  5,
		},
		{
			clothes: [][]string{{"crow_mask", "face"}, {"blue_sunglasses", "face"}, {"smoky_makeup", "face"}},
			result:  3,
		},
		{
			clothes: [][]string{{"crow_mask", "face"}},
			result:  1,
		},
		{
			clothes: [][]string{{"crow_mask", "face"}, {"blue_sunglasses", "face"}, {"yellow_hat", "headgear"}, {"blue_sunglasses", "eyewear"}, {"green_turban", "headgear"}},
			result:  17,
		},
	}

	for i, tc := range testCaseArr {
		if solution(tc.clothes) != tc.result {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

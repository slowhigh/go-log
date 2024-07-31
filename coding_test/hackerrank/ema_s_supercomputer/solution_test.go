package emaSSupercomputer

import "testing"

type testCase struct {
	grid   []string
	result int32
}

func TestTwoPluses(t *testing.T) {
	testCaseArr := []testCase{
		{
			grid: []string{
				"GGGGGGGGGGGG",
				"GGGGGGGGGGGG",
				"BGBGGGBGBGBG",
				"BGBGGGBGBGBG",
				"GGGGGGGGGGGG",
				"GGGGGGGGGGGG",
				"GGGGGGGGGGGG",
				"GGGGGGGGGGGG",
				"BGBGGGBGBGBG",
				"BGBGGGBGBGBG",
				"BGBGGGBGBGBG",
				"BGBGGGBGBGBG",
				"GGGGGGGGGGGG",
				"GGGGGGGGGGGG",
			},
			result: 189,
		},
		{
			grid: []string{
				"BGB",
				"GGG",
				"BGB",
			},
			result: 1,
		},
		{
			grid: []string{
				"BB",
				"BB",
			},
			result: 0,
		},
		{
			grid: []string{
				"GG",
				"GG",
			},
			result: 1,
		},
		{
			grid: []string{
				"GGGGGG",
				"GBBBGB",
				"GGGGGG",
				"GGBBGB",
				"GGGGGG",
			},
			result: 5,
		},
		{
			grid: []string{
				"BGBBGB",
				"GGGGGG",
				"BGBBGB",
				"GGGGGG",
				"BGBBGB",
				"BGBBGB",
			},
			result: 25,
		},
	}

	for i, tc := range testCaseArr {
		res := twoPluses(tc.grid)

		if res == tc.result {
			t.Logf("PASS - %d(%d)", i, res)
		} else {
			t.Errorf("FAIL - %d(%d)", i, res)
		}
	}
}

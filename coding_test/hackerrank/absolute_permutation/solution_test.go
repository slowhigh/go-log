package absolutePermutation

import (
	"fmt"
	"testing"
)

type testCase struct {
	n      int32
	k      int32
	result []int32
}

func Test_absolutePermutation(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:      1,
			k:      0,
			result: []int32{1},
		},
		{
			n:      1,
			k:      1,
			result: []int32{-1},
		},
		{
			n:      6,
			k:      3,
			result: []int32{4, 5, 6, 1, 2, 3},
		},
		{
			n:      8,
			k:      2,
			result: []int32{3, 4, 1, 2, 7, 8, 5, 6},
		},
		{
			n:      8,
			k:      4,
			result: []int32{5, 6, 7, 8, 1, 2, 3, 4},
		},
		{
			n:      4,
			k:      4,
			result: []int32{-1},
		},
		{
			n:      6,
			k:      2,
			result: []int32{-1},
		},
		{
			n:      6,
			k:      4,
			result: []int32{-1},
		},
		{
			n:      6,
			k:      5,
			result: []int32{-1},
		},
		{
			n:      2,
			k:      1,
			result: []int32{2, 1},
		},
		{
			n:      3,
			k:      0,
			result: []int32{1, 2, 3},
		},
		{
			n:      3,
			k:      2,
			result: []int32{-1},
		},
		{
			n:      8,
			k:      5,
			result: []int32{-1},
		},
	}

	for i, tc := range testCaseArr {
		res := fmt.Sprintf("%+v", absolutePermutation(tc.n, tc.k))
		if res == fmt.Sprintf("%+v", tc.result) {
			t.Logf("PASS - %d (%s)", i, res)
		} else {
			t.Errorf("FAIL - %d (%s)", i, res)
		}
	}
}

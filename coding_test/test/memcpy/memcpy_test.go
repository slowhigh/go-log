package memcpy

import (
	"fmt"
	"testing"
)

type testCase struct {
	data   []byte
	src    int64
	dest   int64
	size   int64
	result []byte
}

func Test_memcpy(t *testing.T) {
	testCaseArr := []testCase{
		{
			data:   []byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
			src:    2,
			dest:   4,
			size:   4,
			result: []byte{1, 2, 3, 4, 3, 4, 5, 6, 9},
		},
		{
			data:   []byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
			src:    4,
			dest:   2,
			size:   4,
			result: []byte{1, 2, 5, 6, 7, 8, 7, 8, 9},
		},
	}

	for i, tc := range testCaseArr {
		memcpy(tc.data, tc.src, tc.dest, tc.size)

		if fmt.Sprintf("%+v", tc.data) == fmt.Sprintf("%+v", tc.result) {
			t.Logf("PASS - %d(%+v)", i, tc.data)
		} else {
			t.Logf("FAIL - %d(%+v)", i, tc.data)
		}
	}
}

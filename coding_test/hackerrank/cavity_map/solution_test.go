package cavityMap

import (
	"fmt"
	"testing"
)

type testCase struct {
	grid   []string
	result []string
}

func Test_cavityMap(t *testing.T) {
	testCaseArr := []testCase{
		{
			grid: []string{
				"1112",
				"1912",
				"1892",
				"1234",
			},
			result: []string{
				"1112",
				"1X12",
				"18X2",
				"1234",
			},
		},
	}

	for i, tc := range testCaseArr {
		res := fmt.Sprintf("%+v", cavityMap(tc.grid))
		if res == fmt.Sprintf("%+v", tc.result) {
			t.Logf("PASS - %d (%s)", i, res)
		} else {
			t.Errorf("FAIL - %d (%s)", i, res)
		}
	}
}

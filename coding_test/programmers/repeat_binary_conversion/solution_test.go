package repeat_binary_conversion

import (
	//"fmt"
	"fmt"
	"testing"
)

type testCase struct {
	s      string
	result []int
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			s:      "110010101001",
			result: []int{3, 8},
		},
		{
			s:      "01110",
			result: []int{3, 3},
		},
		{
			s:      "1111111",
			result: []int{4, 1},
		},
	}

	for i, tc := range testCaseArr {
		if fmt.Sprint(solution(tc.s)) != fmt.Sprint(tc.result) {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

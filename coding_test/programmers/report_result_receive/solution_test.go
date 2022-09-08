package report_result_receive

import (
	"fmt"
	"testing"
)

type testCase struct {
	id_list []string
	report  []string
	k       int
	result  []int
}

func Test_solution(t *testing.T) {
	testCaseArr := []testCase{
		{
			id_list: []string{"muzi", "frodo", "apeach", "neo"},
			report:  []string{"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi"},
			k:       2,
			result:  []int{2, 1, 1, 0},
		},
		{
			id_list: []string{"con", "ryan"},
			report:  []string{"ryan con", "ryan con", "ryan con", "ryan con"},
			k:       3,
			result:  []int{0, 0},
		},
		{
			id_list: []string{"a", "b", "c", "d"},
			report:  []string{"a e", "e a", "d c", "b d", "a b"},
			k:       1,
			result:  []int{1, 1, 0, 1},
		},
	}

	for i, tc := range testCaseArr {
		if fmt.Sprint(solution(tc.id_list, tc.report, tc.k)) != fmt.Sprint(tc.result) {
			t.Errorf("Wrong result - %d", i)
		}
	}
}

package report_result_receive

import "strings"

func solution(id_list []string, report []string, k int) []int {
	mailCountArr := make([]int, len(id_list))
	reportersByIdMap := make(map[string][]string)
	reportCountMap := make(map[string]int)

	for _, id := range id_list {
		reportersByIdMap[id] = []string{}
		reportCountMap[id] = 0
	}

	for _, r := range report {
		s := strings.Split(r, " ")
		if _, ok := reportersByIdMap[s[0]]; !ok {
			continue
		}
		if _, ok := reportersByIdMap[s[1]]; !ok {
			continue
		}

		isContain := false
		for _, reporter := range reportersByIdMap[s[1]] {
			if reporter == s[0] {
				isContain = true
				break
			}
		}
		if !isContain {
			reportersByIdMap[s[1]] = append(reportersByIdMap[s[1]], s[0])
		}
	}

	for id := range reportersByIdMap {
		if len(reportersByIdMap[id]) < k {
			continue
		}
		for _, reporter := range reportersByIdMap[id] {
			reportCountMap[reporter] += 1
		}
	}

	for i := 0; i < len(mailCountArr); i++ {
		mailCountArr[i] = reportCountMap[id_list[i]]
	}

	return mailCountArr
}

package main

import (
	"fmt"
	"strings"
)

func solution(id_list []string, report []string, k int) []int {
	id_map := make(map[string]string)
	result_map := make(map[string]int)
	result_arr := make([]int, len(id_list))

	for i := 0; i < len(id_list); i++ {
		id_map[id_list[i]] = ""
		result_map[id_list[i]] = 0
	}

	for i := 0; i < len(report); i++ {
		report_split_array := strings.Split(report[i], " ")

		if strings.Contains(id_map[report_split_array[1]], report_split_array[0]) {
			continue
		}

		id_map[report_split_array[1]] += report_split_array[0] + " "
	}

	for _, value := range id_map {
		value_split := strings.Split(strings.TrimSpace(value), " ")

		if len(value_split) < k {
			continue
		}

		for j := 0; j < len(value_split); j++ {
			result_map[value_split[j]] += 1
		}
	}

	for i, v := range id_list {
		result_arr[i] = result_map[v]
	}

	return result_arr
}

func main() {
	id_list := []string{"muzi", "frodo", "apeach", "neo"}
	report := []string{"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi"}
	k := 2
	fmt.Println(solution(id_list, report, k))
}

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

		if _, ok := id_map[report_split_array[0]]; !ok {
			continue
		}

		if _, ok := id_map[report_split_array[1]]; !ok {
			continue
		}

		if strings.Contains(id_map[report_split_array[1]], report_split_array[0]) {
			continue
		}

		id_map[report_split_array[1]] += report_split_array[0] + " "

		fmt.Printf("id_map: %v\n", id_map)
	}

	for _, value := range id_map {
		value_split := strings.Split(strings.TrimSpace(value), " ")

		if len(value_split) < k {
			continue
		}

		for j := 0; j < len(value_split); j++ {
			result_map[value_split[j]] = result_map[value_split[j]] + 1
		}
	}

	for i, v := range id_list {
		result_arr[i] = result_map[v]
	}

	return result_arr
}

// 2 ≤ id_list의 길이 ≤ 1,000
// 1 ≤ id_list의 원소 길이 ≤ 10
// id_list의 원소는 이용자의 id를 나타내는 문자열이며 알파벳 소문자로만 이루어져 있습니다.
// id_list에는 같은 아이디가 중복해서 들어있지 않습니다.
// 1 ≤ report의 길이 ≤ 200,000
// 3 ≤ report의 원소 길이 ≤ 21
// report의 원소는 "이용자id 신고한id"형태의 문자열입니다.
// 예를 들어 "muzi frodo"의 경우 "muzi"가 "frodo"를 신고했다는 의미입니다.
// id는 알파벳 소문자로만 이루어져 있습니다.
// 이용자id와 신고한id는 공백(스페이스)하나로 구분되어 있습니다.
// 자기 자신을 신고하는 경우는 없습니다.
// 1 ≤ k ≤ 200, k는 자연수입니다.
// return 하는 배열은 id_list에 담긴 id 순서대로 각 유저가 받은 결과 메일 수를 담으면 됩니다.

func main() {
	id_list := []string{"a", "b", "c", "d"}
	report := []string{"a e", "e a", "d c", "b d", "a b"}
	k := 2
	fmt.Println(solution(id_list, report, k))
}

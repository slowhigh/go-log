package main

import (
	"fmt"
	"strings"
)

func solution(id_list []string, report []string, k int) []int {

	id_map := make(map[string][]string)
	for i := 0; i < len(id_list); i++ {
		id_map[id_list[i]] = make([]string, 0, 4)
	}

	for _, item := range report {
		item_array := strings.Split(item, " ")

		strings.Contains(id_map[item_array[]])
		id_map[item_array[0]] = append(id_map[item_array[0]], item_array[1])
	}

	return []int{}
}

func main() {
	id_list := []string{"muzi", "frodo", "apeach", "neo"}
	report := []string{"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi"}
	k := 2
	fmt.Println(solution(id_list, report, k))
}

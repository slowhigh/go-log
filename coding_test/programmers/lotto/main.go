package main

//max, min := 0, 0 변수 한줄 초기화

import (
	"fmt"
)

func solution(lottos []int, win_nums []int) []int {
	max, min := 0, 0

	for _, lotto := range lottos {
		if lotto == 0 {
			max++
			continue
		}

		var index int = -1
		for i, win_num := range win_nums {
			if lotto == win_num {
				max++
				min++
				index = i
				break
			}
		}

		if index != -1 {
			win_nums = append(win_nums[:index], win_nums[index+1:]...)
		}
	}

	if 7-max == 7 {
		max = 1
	}

	if 7-min == 7 {
		min = 1
	}

	return []int{7 - max, 7 - min}
}

func main() {

	a := []int{44, 1, 0, 0, 31, 25}
	b := []int{31, 10, 45, 1, 6, 19}

	fmt.Println(solution(a, b))
}

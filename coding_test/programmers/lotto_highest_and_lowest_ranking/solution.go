package lotto_highest_and_lowest_ranking

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

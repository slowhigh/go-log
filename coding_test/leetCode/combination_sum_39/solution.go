package combination_sum_39

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	combiArr := make([][]int, 0)

	sort.Ints(candidates)
	for i := 0; i < len(candidates); i++ {
		remain := target - candidates[i]

		if remain < 0 {
			continue
		} else if remain == 0 {
			combiArr = append(combiArr, []int{candidates[i]})
			continue
		}

		subCombiArr := combinationSum(candidates[i:], remain)

		for j := 0; j < len(subCombiArr); j++ {
			combiArr = append(combiArr, append([]int{candidates[i]}, subCombiArr[j]...))
		}
	}

	return combiArr
}

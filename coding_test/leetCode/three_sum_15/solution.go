package three_sum_15

import "sort"

type threeSumItem struct {
	nums []int
	len  int
	sum  int
}

func threeSum(nums []int) [][]int {
	threeSumItemMap := make(map[int]threeSumItem)
	result := [][]int{}

	for _, num := range nums {
		keys := make([]int, 0, len(threeSumItemMap))
		for k := range threeSumItemMap {
			keys = append(keys, k)
		}

		for key := range keys {
			newThreeSumItem := threeSumItem{}
			newThreeSumItem.nums = append(threeSumItemMap[key].nums, num)
			newThreeSumItem.len = threeSumItemMap[key].len + 1
			newThreeSumItem.sum = threeSumItemMap[key].sum + num

			if newThreeSumItem.len < 3 {
				threeSumItemMap[len(threeSumItemMap)] = newThreeSumItem
				continue
			}

			if newThreeSumItem.sum == 0 {
				sort.Ints(newThreeSumItem.nums)
				result = append(result, newThreeSumItem.nums)
			}
		}

		threeSumItemMap[len(threeSumItemMap)] = threeSumItem{nums: []int{num}, len: 1, sum: num}
	}

	return result
}

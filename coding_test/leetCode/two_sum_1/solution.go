package two_sum_1

import (
	"sort"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		rest := target - nums[i];
		index := sort.SearchInts(nums[i+1:], rest)

		if rest == nums[i+1+index] {
			return []int {i, i+1+index}
		}
	}

	return []int {}
}

func twoSum2(nums []int, target int) []int {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		rest := target - nums[i];

		for j := i + 1; j < numsLen; j++ {
			if nums[j] == rest {
				return []int {i, j}
			}
		}
	}

	return []int {}
}
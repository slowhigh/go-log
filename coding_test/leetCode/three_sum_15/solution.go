package three_sum_15

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	result := [][]int{}
	sort.Ints(nums)

	if numsLen < 3 {
		return [][]int{}
	} else if numsLen == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			return [][]int{{nums[0], nums[1], nums[2]}}
		}
		return [][]int{}
	}

	for pivot := 0; pivot < numsLen-2; pivot++ {
		if pivot > 0 && nums[pivot] == nums[pivot-1] {
			continue
		}

		left, right := pivot+1, numsLen-1

		for left < right {
			if left > pivot+1 && nums[left] == nums[left-1] {
				left++
				continue
			} else if right < numsLen-1 && nums[right] == nums[right+1] {
				right--
				continue
			}

			sum := nums[pivot] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[pivot], nums[left], nums[right]})
				left++
				right--
			} else if sum < 0 {
				left++
			} else { // sum > 0
				right--
			}
		}
	}

	return result
}

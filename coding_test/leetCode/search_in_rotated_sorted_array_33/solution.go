package search_in_rotated_sorted_array_33

func search(nums []int, target int) int {
	numsLen := len(nums)

	if numsLen == 1 && nums[0] == target {
		return 0
	}

	left, right := 0, numsLen-1

	for left <= right {
		if right-left < 2 {
			if nums[right] == target {
				return right
			} else if nums[left] == target {
				return left
			} else {
				break
			}
		}

		mid := (right-left)/2 + left

		if (nums[left] <= nums[mid] && (nums[left] <= target && target <= nums[mid])) || (nums[mid+1] <= nums[right] && (target < nums[mid+1] || nums[right] < target)) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return -1
}

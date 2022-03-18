package find_minimum_in_rotated_sorted_array_153

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (right - left) / 2 + left
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
		if nums[left] < nums[right] {
			return nums[left]
		}
	}

	return nums[left]
}

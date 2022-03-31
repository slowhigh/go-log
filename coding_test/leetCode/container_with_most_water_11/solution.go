package container_with_most_water_11

func maxArea(height []int) int {
    res := 0
	left, right := 0, len(height) - 1

	for left < right {
		area := 0
		if height[left] < height[right] {
			area = (right - left) * height[left]
			left++
		} else {
			area = (right - left) * height[right]
			right--
		}

		if area > res {
			res = area
		}
	}

	return res
}
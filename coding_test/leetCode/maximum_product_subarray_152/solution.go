package maximum_product_subarray_152

func maxProduct1(nums []int) int {
	numsLen := len(nums)
	maxPro := nums[0]

	for i := 0; i < numsLen; i++ {
		curPro := 1

		for j := i; j < numsLen; j++ {
			curPro *= nums[j]

			if maxPro < curPro {
				maxPro = curPro
			}
		}
	}

	return maxPro
}

func maxProduct2(nums []int) int {
	res, curMin, curMax := nums[0], nums[0], nums[0]

	for _, num := range nums[1:] {
		curMin, curMax = curMin * num, curMax * num
		curMin, curMax = min(min(curMax, num), curMin), max(max(curMax, num), curMin)

		if res < curMax {
			res = curMax
		}
	}

	return res
}

func max(x int, y int) int {
	if x < y { return y }
	return x
}

func min(x int, y int) int {
	if x < y { return x }
	return y
}
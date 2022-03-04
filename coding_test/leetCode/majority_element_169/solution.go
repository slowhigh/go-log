package majority_element_169

func majorityElement(nums []int) int {
	checkMap := make(map[int]int)
	maxIndex, maxValue := 0, 0

	for i := 0; i < len(nums); i++ {
		checkMap[nums[i]]++
	}

	for k, v := range checkMap {
		if maxValue <= v {
			maxIndex = k
			maxValue = v
		}
	}

	return maxIndex
}

func majorityElement2(nums []int) int {
	cnt := 0
	result := 0
	for i := 0; i < len(nums); i++ {
		if cnt == 0 {
			result = nums[i]
			cnt++
		} else if nums[i] == result {
			cnt++
		} else {
			cnt--
		}
	}
	return result
}

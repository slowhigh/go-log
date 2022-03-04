package subarray_sum_equals_k_560

import "fmt"

func subarraySum(nums []int, k int) int {
	equalsCount := 0
	numsLen := len(nums)
	subNumsSum := 0
	lastNumIndex := 0

	// 서브 배열의 길이
	for i := 1; i <= numsLen; i++ {
		subNumsSum = 0

		// 서브 배열의 첫 index
		for j := 0; j < numsLen; j++ {
			lastNumIndex = j + i

			if lastNumIndex > numsLen {
				break
			}

			if j == 0 {
				for _, v := range nums[j:lastNumIndex] {
					subNumsSum += v
				}
			} else {
				subNumsSum = subNumsSum - nums[j-1] + nums[lastNumIndex-1]
			}

			if subNumsSum != k {
				continue
			}

			equalsCount++
		}
	}

	return equalsCount
}

func subarraySum2(nums []int, k int) int {
	dict := make(map[int]int)
	dict[0] = 1
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		fmt.Printf("sum: %v\n", sum)
		res += dict[sum-k]
		fmt.Printf("res: %v\n", res)
		fmt.Printf("dict: %v\n", dict)
		dict[sum]++
		fmt.Printf("dict: %v\n", dict)
	}
	return res
}

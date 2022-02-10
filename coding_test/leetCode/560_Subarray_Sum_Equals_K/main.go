package main

import "fmt"

// Given an array of integers nums and an integer k, return the total number of continuous subarrays whose sum equals to k.

// Example 1:

// Input: nums = [1,1,1], k = 2
// Output: 2
// Example 2:

// Input: nums = [1,2,3], k = 3
// Output: 2

// Constraints:

// 1 <= nums.length <= 2 * 104
// -1000 <= nums[i] <= 1000
// -107 <= k <= 107

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
        res += dict[sum - k]
		fmt.Printf("res: %v\n", res)
		fmt.Printf("dict: %v\n", dict)
        dict[sum]++
		fmt.Printf("dict: %v\n", dict)
    }
    return res
}

func main() {
	nums := []int{1,2,3}
	k := 3

	fmt.Println(subarraySum(nums, k))
	fmt.Println(subarraySum2(nums, k))
}

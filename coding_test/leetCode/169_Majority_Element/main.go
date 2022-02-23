package main

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

// Example 1:
// Input: nums = [3,2,3]
// Output: 3

// Example 2:
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

// Constraints:
// n == nums.length
// 1 <= n <= 5 * 10^4
// -2^31 <= nums[i] <= 2^31 - 1

// Follow-up: Could you solve the problem in linear time and in O(1) space?

import (
	"fmt"
	"math"
)

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

func main() {

	//nums := []int{3, 2, 3}
	//nums2 := []int{2, 2, 1, 1, 1, 2, 2}
	nums3 := []int{(int(math.Pow(2, 31)) * -1), 1, 1, 1, 2, 1}

	//fmt.Println(majorityElement(nums))
	//fmt.Println(majorityElement(nums2))
	fmt.Println(majorityElement(nums3))
}

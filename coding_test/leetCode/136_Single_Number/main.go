package main

import (
	"fmt"
	"sort"
)

// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

// You must implement a solution with a linear runtime complexity and use only constant extra space.

 
// Example 1:
// Input: nums = [2,2,1]
// Output: 1

// Example 2:
// Input: nums = [4,1,2,1,2]
// Output: 4

// Example 3:
// Input: nums = [1]
// Output: 1
 

// Constraints:
// 1 <= nums.length <= 3 * 10^4
// -3 * 10^4 <= nums[i] <= 3 * 10^4

// Each element in the array appears twice except for one element which appears only once.

func singleNumber(nums []int) (single int) {
	checkMap := make(map[int] struct{})

	for i := 0; i < len(nums); i++ {
		if _, exists := checkMap[nums[i]]; exists {
			delete(checkMap, nums[i])
		} else {
			checkMap[nums[i]] = struct{}{}
		}
	}

	for k := range checkMap {
		single = k
	}

	return
}

func singleNumber2(nums []int) int {
	sum := 0;
	_nums := make([]int, len(nums))
	copy(_nums, nums)
	sort.Ints(_nums)

	for i := 0; i < len(_nums); i++ {
		if i % 2 == 0 {
			sum += _nums[i]
		} else {
			sum -= _nums[i]
		}
	}

	return sum
}

func singleNumber3(nums []int) int {
	sum := 0;

	for i := 0; i < len(nums); i++ {
		fmt.Printf("nums[i]: %v\n", nums[i])
		sum ^= nums[i]
		fmt.Printf("sum: %v\n", sum)
	}

	return sum
}

func main() {
	nums := []int {2,3,2,3,1}
	fmt.Println(singleNumber(nums))
	fmt.Println(singleNumber2(nums))
	fmt.Println(singleNumber3(nums))
}
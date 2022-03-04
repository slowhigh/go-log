package single_number_136

import (
	"fmt"
	"sort"
)

func singleNumber(nums []int) (single int) {
	checkMap := make(map[int]struct{})

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
	sum := 0
	_nums := make([]int, len(nums))
	copy(_nums, nums)
	sort.Ints(_nums)

	for i := 0; i < len(_nums); i++ {
		if i%2 == 0 {
			sum += _nums[i]
		} else {
			sum -= _nums[i]
		}
	}

	return sum
}

func singleNumber3(nums []int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		fmt.Printf("nums[i]: %v\n", nums[i])
		sum ^= nums[i]
		fmt.Printf("sum: %v\n", sum)
	}

	return sum
}

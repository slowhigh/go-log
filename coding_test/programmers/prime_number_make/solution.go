package prime_number_make

import (
	"sort"
)

func solution(nums []int) int {
	primeNumberCount := 0
	stack := make([]int, 0)

	sort.Ints(nums)
	numsLen := len(nums)

	for i := 0; ; i++ {
		i %= numsLen
		stackLen := len(stack)

		if stackLen == 0 {
			stack = append(stack, nums[i])
			continue
		}

		if stack[stackLen-1] > nums[i] {
			continue
		}

		if stack[stackLen-1] == nums[i] {
			stack = stack[:stackLen-1]
			continue
		}

		if stack[0] == nums[len(nums)-2] {
			break
		}

		stack = append(stack, nums[i])
		stackLen = len(stack)

		if stackLen == 3 {
			sum := stack[0] + stack[1] + stack[2]

			if checkPrimeNumber(sum) {
				primeNumberCount++
			}

			stack = stack[:stackLen-1]
			continue
		}
	}

	return primeNumberCount
}

func checkPrimeNumber(value int) bool {
	for i := 2; i < value; i++ {
		if value%i == 0 {
			return false
		}
	}

	return true
}

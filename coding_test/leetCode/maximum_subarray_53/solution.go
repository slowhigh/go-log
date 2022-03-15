package maximum_subarray_53

import "math"

func maxSubArray(nums []int) int {
    max := 0
    
    var maxSum func(subArr []int, isLeft bool) int

    maxSum = func(subArr []int, isLeft bool) int {
        subArrLen := len(subArr)

        if subArrLen == 1 {
            return subArr[0]
        }

        mid := int(subArrLen / 2)

        leftSum := maxSum(subArr[:mid], true)
        rightSum := maxSum(subArr[mid + 1:], false)
        totalSum := leftSum + rightSum

        subMax := int(math.Max(math.Max(float64(leftSum), float64(rightSum)), float64(totalSum)))

        if max < subMax {
            max = subMax
        }

        if isLeft && {
            return rightSum
        }
    }

    maxSum(nums, true)
    
    
    
    
    return max
}
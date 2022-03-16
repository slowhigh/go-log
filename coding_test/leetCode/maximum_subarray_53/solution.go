package maximum_subarray_53

func maxSubArray(nums []int) int {
    max, sum := nums[0], nums[0]
    
    for i := 1; i < len(nums); i++ {
        sum += nums[i]
        if sum < nums[i] {
            sum = nums[i]
        }

        if max < sum {
            max = sum
        }
    }

    return max
}
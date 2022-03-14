package product_of_array_except_self_238

func productExceptSelf(nums []int) []int {
    numsLen := len(nums)
    pro := 1
    proArr := make([]int, numsLen)
    proArr[0] = 1
    
    for i := 1; i < numsLen; i++ {
        pro *= nums[i - 1]
        proArr[i] = pro
    }
    
    pro = 1
    
    for i := numsLen - 2; i > -1; i-- {
        pro *= nums[i + 1]
        proArr[i] *= pro
    }
    
    return proArr
}
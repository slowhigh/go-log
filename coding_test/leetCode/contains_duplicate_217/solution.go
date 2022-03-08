package contains_duplicate_217

func containsDuplicate(nums []int) bool {
    duplicationMap := make(map[int]struct{})
    
    for i := 0; i < len(nums); i++ {
        if _, ok := duplicationMap[nums[i]]; ok {
            return true
        }
        
        duplicationMap[nums[i]] = struct{}{}
    }
    
    return false
}
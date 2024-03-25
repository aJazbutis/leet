func findDuplicates(nums []int) []int {
    ret := []int{}
    for i := range nums {
        j := int(math.Abs(float64(nums[i])))-1
        if nums[j] > 0 {
            nums[j] = -nums[j]
        } else {
            ret = append(ret, j+1)
        }
    }
    return ret
}

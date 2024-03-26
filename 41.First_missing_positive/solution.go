func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := range nums {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = n + 1
		}
	}

	for i := range nums {
		if int(math.Abs(float64(nums[i]))) < n+1 {
			if nums[int(math.Abs(float64(nums[i])))-1] > 0 {
				nums[int(math.Abs(float64(nums[i])))-1] *= -1
			}
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}


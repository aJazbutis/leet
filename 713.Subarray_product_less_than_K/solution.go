func numSubarrayProductLessThanK(nums []int, k int) int {
	l, r, c, product := 0, 0, 0, 1
	n := len(nums)
	if k <= 1 {
		return 0
	}
	for r < n {
		product *= nums[r]
		for product >= k {
			product /= nums[l]
			l++
		}
		c += 1 + (r - l)
		r++
	}
	return c
}

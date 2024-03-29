func countSubarrays(nums []int, k int) int64 {
	l, r, c, kc := 0, 0, 0, 0
	max := slices.Max(nums)
	n := len(nums)
	for r < n {
		if nums[r] == max {
			kc++
			if kc == k {
				c += n - r
				for nums[l] != max {
					l++
					c += n - r
				}
				l++
				kc--
			}
		}
		r++
	}
	return int64(c)
}

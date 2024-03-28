func maxSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 0
	}
	l, r, ret := 0, 0, 0
	n := len(nums)
	freq := make(map[int]int)
	for r < n {
		freq[nums[r]] += 1
		for freq[nums[r]] > k {
			freq[nums[l]] -= 1
			l++
		}
		ret = max(ret, r-l+1)
		r++
	}
	return ret
}

func maximumHappinessSum(happiness []int, k int) int64 {
	var sum int64
	var idx int
	slices.SortFunc(happiness, func(a, b int) int {
		return b - a
	})
	for k > 0 {
		h := int64(happiness[idx] - idx)
		if h < 1 {
			break
		}
		sum += h
		idx++
		k--
	}
	return sum
}

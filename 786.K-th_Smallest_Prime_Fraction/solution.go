func kthSmallestPrimeFraction(arr []int, k int) []int {
	fractions := [][]int{}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			fractions = append(fractions, []int{arr[i], arr[j]})
		}
	}
	slices.SortFunc(fractions, func(a, b []int) int {
		x, y := float64(a[0])/float64(a[1]), float64(b[0])/float64(b[1])
		if x >= y {
			return 1
		}
		return -1
	})
	return fractions[k-1]
}

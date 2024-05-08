func findRelativeRanks(score []int) []string {
	ret := make([]string, len(score))
	sorted := make([][2]int, len(score))
	for i := range sorted {
		sorted[i] = [2]int{i, score[i]}
	}
	slices.SortFunc(sorted, func(a, b [2]int) int {
		return b[1] - a[1]
	})
	for i, v := range sorted {
		var place string
		switch i {
		case 0:
			place = "Gold Medal"
		case 1:
			place = "Silver Medal"
		case 2:
			place = "Bronze Medal"
		default:
			place = strconv.Itoa(i + 1)
		}
		ret[v[0]] = place
	}
	return ret
}

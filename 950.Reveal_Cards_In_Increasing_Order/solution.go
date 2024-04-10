func deckRevealedIncreasing(deck []int) []int {
	slices.Sort(deck)
	ret := make([]int, len(deck))
	idxs := make([]int, len(deck))
	for i := range idxs {
		idxs[i] = i
	}
	for _, v := range deck {
		idx := idxs[0]
		ret[idx] = v
		idxs = idxs[1:]
		if len(idxs) > 1 {
			idxs = append(idxs[1:], idxs[0])
		}
	}
	return ret
}

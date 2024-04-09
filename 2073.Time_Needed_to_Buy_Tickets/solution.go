func timeRequiredToBuy(tickets []int, k int) int {
	c := 0
	for {
		for i := range tickets {
			if tickets[i] == 0 {
				continue
			}
			c++
			tickets[i]--
			if i == k && tickets[i] == 0 {
				return c
			}
		}
	}
	return c
}

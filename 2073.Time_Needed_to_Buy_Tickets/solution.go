func timeRequiredToBuy(tickets []int, k int) int {
	c := 0
	for i := range tickets {
		if i <= k {
			if tickets[i] <= tickets[k] {
				c += tickets[i]
			} else {
				c += tickets[k]
			}
		} else {
			if tickets[i] < tickets[k] {
				c += tickets[i]
			} else {
				c += tickets[k] - 1
			}
		}
	}
	return c
}

//brute as always inital :D
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

func numRescueBoats(people []int, limit int) int {
	c, left := 0, len(people)
	slices.Sort(people)
	i, j := 0, len(people)-1
	for i < j {
		if people[i]+people[j] <= limit {
			i++
			j--
			c++
			left -= 2
		} else {
			c++
			j--
			left--
		}
	}
	if left != 0 {
		c++
	}
	return c
}

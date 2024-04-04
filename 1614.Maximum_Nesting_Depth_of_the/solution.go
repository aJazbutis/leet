func maxDepth(s string) int {
	maxDepth := 0
	curDepth := 0
	for _, v := range s {
		switch v {
		case '(':
			curDepth++
		case ')':
			maxDepth = max(maxDepth, curDepth)
			curDepth--
		}
	}
	return maxDepth
}

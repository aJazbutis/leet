func leftMax(height []int) int {
	m := height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		m = max(m, height[i])
	}
	return m
}

func rightMax(height []int) int {
	m := height[0]
	for i := 1; i < len(height); i++ {
		m = max(m, height[i])
	}
	return m
}

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	ret := 0
	m := -1
	lMax, rMax := 0, 0
	for i := 1; i < len(height)-1; i++ {
		if m == -1 {
			lMax = leftMax(height[:i])
			rMax = rightMax(height[i+1:])
			m = min(lMax, rMax)
		}
		if m > height[i] {
			ret += m - height[i]
		} else {
			lMax = height[i]
			rMax = rightMax(height[i+1:])
			m = min(lMax, rMax)
		}
	}
	return ret
}

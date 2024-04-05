//the good one
func makeGood(s string) string {
    stack := []rune{}
    for i := range s {
        if len(stack) > 0 && int(math.Abs(float64(stack[len(stack)-1] - rune(s[i])))) == 32 {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, rune(s[i]))
        }
    }
	return string(stack)
}

// the first instinct ;(
func makeGood(s string) string {
	for i := 0; i < len(s)-1; i++ {
		if int(math.Abs(float64(rune(s[i])-rune(s[i+1])))) == 32 {
			s = s[:i] + s[i+2:]
			i -= 2
			if i < -1 {
				i = -1
			}
		}
	}
	return s
}

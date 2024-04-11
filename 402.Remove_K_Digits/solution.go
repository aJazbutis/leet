//monotonic stack
func removeKdigits(num string, k int) string {
	if k == 0 {
		return num
	}
	stack := []rune{}
	for _, v := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > v {
			k--
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}
    stack = stack[:len(stack)-k]
	k = 0
	for _, v := range stack {
		if v == '0' {
			k++
		} else {
			break
		}
	}
	stack = stack[k:]
	if len(stack) == 0 {
		return "0"
	} else {
		return string(stack)
	}
}

//first brute
func removeKdigits(num string, k int) string {
	if k == 0 {
		return num
	}
	stack := []rune(num)
	for k > 0 {
		l := len(stack)
		for i := 0; i < len(stack)-1; i++ {
			if stack[i] > stack[i+1] {
				stack = append(stack[:i], stack[i+1:]...)
				break
			}
		}
		if len(stack) == l {
			stack = stack[:len(stack)-k]
			k = 0
		}
		k--
	}
	k = 0
	for _, v := range stack {
		if v == '0' {
			k++
		} else {
			break
		}
	}
	stack = stack[k:]
	if len(stack) == 0 {
		return "0"
	} else {
		return string(stack)
	}
}

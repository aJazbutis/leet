//the good one
func minRemoveToMakeValid(s string) string {
	stack := []int{}
	invalid := make([]bool, len(s))
	for i, v := range s {
		switch v {
		case '(':
			stack = append(stack, i)
		case ')':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				invalid[i] = true
			}
		}
	}
	for i := range stack {
		invalid[stack[i]] = true
	}
	var ret strings.Builder
	for i, v := range s {
		if !invalid[i] {
			ret.WriteRune(v)
		}
	}
	return ret.String()
}

//first
func minRemoveToMakeValid(s string) string {
	c := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			c++
		case ')':
			c--
			if c < 0 {
				s = s[:i] + s[i+1:]
				c = 0
				i--
			}
		}
	}
	c = 0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '(':
			c--
			if c < 0 {
				s = s[:i] + s[i+1:]
				c = 0
			}
		case ')':
			c++
		}
	}
	return s
}

//less reslicing 
func minRemoveToMakeValid(s string) string {
	ss := []rune(s)
	c := 0
	cc := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			c++
		case ')':
			c--
			if c < 0 {
				ss[i] = '-'
				c = 0
				cc++
			}
		}
	}
	c = 0
	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '(':
			c--
			if c < 0 {
				ss[i] = '-'
				c = 0
				cc++
			}
		case ')':
			c++
		}
	}
	ret := make([]rune, len(s)-cc)
	for i, j := 0, 0; i < len(ss); i++ {
		if ss[i] != '-' {
			ret[j] = ss[i]
			j++
		}
	}
	return string(ret)
}

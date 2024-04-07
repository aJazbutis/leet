func checkValidString(s string) bool {
	open := 0
    for _, v := range s {
		switch v {
		case '(', '*':
            open++
		case ')':
            switch {
                case open == 0:
                    return false
                default:
                    open--
            }
        }
    }
    close := 0
    for i := len(s)-1; i >= 0; i-- {
		switch s[i] {
		case ')', '*':
            close++
		case '(':
            switch {
                case close == 0:
                    return false
                default:
                    close--
            }
        }
    }
    return true
}

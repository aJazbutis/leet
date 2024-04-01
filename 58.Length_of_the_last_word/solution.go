func lengthOfLastWord(s string) int {
    l := 0;
    s = strings.Trim(s, " ")
    for i := len(s)-1; i >=0; i-- {
        if s[i] == ' ' {
            break
        }
        l++
    }
    return l
}
/*or*/
func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	ss := strings.Split(s, " ")
	if len(ss) > 0 {
		return len(ss[len(ss)-1])
	}
	return 0
}

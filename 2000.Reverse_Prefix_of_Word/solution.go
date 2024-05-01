func reversePrefix(word string, ch byte) string {
	idx := strings.IndexByte(word, ch)
	if idx == -1 {
		return word
	}
	// runes := []rune(word)
	// for i := 0; i < idx; i, idx = i+1, idx-1 {
	// runes[i], runes[idx] = runes[idx], runes[i]
	// }
	// return string(runes)
	var b strings.Builder
	for i := idx; i >= 0; i-- {
		b.WriteByte(word[i])
	}
	b.WriteString(word[idx+1:])
	return b.String()
}

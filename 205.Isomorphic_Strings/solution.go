func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mst := make(map[byte]byte)
	mts := make(map[byte]byte)
	for i := range s {
		v, ok := mst[s[i]]
		if ok && v != t[i] {
			return false
		}
		v, ok = mts[t[i]]
		if ok && v != s[i] {
			return false
		}
		mst[s[i]] = t[i]
		mts[t[i]] = s[i]
	}
	return true
}

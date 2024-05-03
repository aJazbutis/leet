func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	var i, j int
	for i < len(v1) && j < len(v2) {
		n1, _ := strconv.Atoi(v1[i])
		n2, _ := strconv.Atoi(v2[j])
		if n1 > n2 {
			return 1
		} else if n2 > n1 {
			return -1
		}
		i++
		j++
	}
	for i < len(v1) {
		n1, _ := strconv.Atoi(v1[i])
		if n1 > 0 {
			return 1
		}
		i++
	}
	for j < len(v2) {
		n2, _ := strconv.Atoi(v2[j])
		if n2 > 0 {
			return -1
		}
		j++
	}
	return 0
}

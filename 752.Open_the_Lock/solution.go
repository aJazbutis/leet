func openLock(deadends []string, target string) int {
	for _, d := range deadends {
		if d == "0000" {
			return -1
		}
	}
	dead, visited1, visited2 := make(map[string]bool), make(map[string]bool), make(map[string]bool)
	for _, v := range deadends {
		dead[v] = true
	}
	queue1, queue2 := []string{"0000"}, []string{target}
	c := -1
	for len(queue1) > 0 && len(queue2) > 0 {
		if len(queue1) > len(queue2) {
			queue1, queue2 = queue2, queue1
			visited1, visited2 = visited2, visited1
		}
		round := len(queue1)
		for round > 0 {
			curr := queue1[0]
			if visited2[curr] {
				return c
			}
			queue1 = queue1[1:]
			visited1[curr] = true
			for i := 0; i < 4; i++ {
				comb := []rune(curr)
				for _, turn := range []int{-1, 1} {
					next := (int(comb[i]-'0')+turn+10)%10 + '0'
					comb[i] = rune(next)
					code := string(comb)
					if !dead[code] && !visited1[code] {
						queue1 = append(queue1, code)
					}
					comb[i] = rune(curr[i])
				}
			}
			round--
		}
		c++
	}
	return -1
}

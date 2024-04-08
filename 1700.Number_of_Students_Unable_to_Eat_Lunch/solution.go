func countStudents(students []int, sandwiches []int) int {
	stud := [2]int{}
	for _, s := range students {
		stud[s]++
	}
	for _, s := range sandwiches {
		if stud[s] > 0 {
			stud[s]--
		} else {
			break
		}
	}
	return stud[0] + stud[1]
}

//;(
func countStudents(students []int, sandwiches []int) int {
	s1 := 0
	s0 := 0
	for _, s := range students {
		switch s {
		case 0:
			s0++
		case 1:
			s1++
		}
	}
	for len(sandwiches) != 0 {
		switch sandwiches[0] {
		case 0:
			if s0 > 0 {
				s0--
				sandwiches = sandwiches[1:]
			} else {
				return s1
			}
		case 1:
			if s1 > 0 {
				s1--
				sandwiches = sandwiches[1:]
			} else {
				return s0
			}
		}
	}
	return 0
}

func tribonacci(n int) int {
	// if n == 0 {
	// 	return 0
	// }
	// if n <= 2 {
	// 	return 1
	// }
	// return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	t0 := 0
	t1 := 1
	t2 := 1
	var curr int
	for i := 3; i <= n; i++ {
		curr = t0 + t1 + t2
		t0, t1, t2 = t1, t2, curr
	}
	return curr
}

func largestLocal(grid [][]int) [][]int {
	ret := make([][]int, len(grid)-2)
	for i := range ret {
		ret[i] = make([]int, len(grid)-2)
	}
	around := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {0, -1}, {1, 1}, {1, 0}, {1, -1}}
	for i := 0; i+1 < len(grid)-1; i++ {
		for j := 0; j+1 < len(grid)-1; j++ {
			y, x := i+1, j+1
			m := grid[y][x]
			for _, c := range around {
				y1 := y + c[0]
				x1 := x + c[1]
				m = max(m, grid[y1][x1])
			}
			ret[i][j] = m
		}
	}
	return ret
}

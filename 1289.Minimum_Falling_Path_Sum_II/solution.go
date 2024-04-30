func minFallingPathSum(grid [][]int) int {
	for y := 1; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			m := int(^uint(0) >> 1)
			for i := 0; i < len(grid[y]); i++ {
				m = min(m, grid[y-1][i])
			}
			if m < grid[y-1][x] {
				grid[y][x] += m
			} else {
				mm := int(^uint(0) >> 1)
				for i := 0; i < len(grid[y]); i++ {
					if x != i {
						mm = min(mm, grid[y-1][i])
					}
				}
                grid[y][x] += mm
			}
		}
	}
	m := int(^uint(0) >> 1)
	for _, v := range grid[len(grid)-1] {
		m = min(m, v)
	}
	return m
}

func islandPerimeter(grid [][]int) int {
	ret := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				ret += 4
				if j-1 >= 0 && grid[i][j-1] == 1 {
					ret -= 2
				}
				if i-1 >= 0 && grid[i-1][j] == 1 {
					ret -= 2
				}
				/*
					if j-1 < 0 || grid[i][j-1] == 0 {
						ret++
					}
					if j+1 == len(grid[i]) || grid[i][j+1] == 0 {
						ret++
					}
					if i-1 < 0 || grid[i-1][j] == 0 {
						ret++
					}
					if i+1 == len(grid) || grid[i+1][j] == 0 {
						ret++
					}
				*/
			}
		}
	}
	return ret
}

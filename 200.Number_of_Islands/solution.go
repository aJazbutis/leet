type Point struct {
	Y, X int
}

func bfs(grid [][]byte, visited [][]bool, dir [4][2]int, start Point) {
	queue := []Point{start}
	visited[start.Y][start.X] = true
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, d := range dir {
			x, y := curr.X+d[0], curr.Y+d[1]
			if (x >= 0 && x < len(grid[0])) && (y >= 0 && y < len(grid)) && !visited[y][x] && grid[y][x] == '1' {
				queue = append(queue, Point{X: x, Y: y})
				visited[y][x] = true
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	dirrections := [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	ret := 0
	for y := range grid {
		for x := range grid[0] {
			if grid[y][x] == '1' && !visited[y][x] {
				bfs(grid, visited, dirrections, Point{Y: y, X: x})
				ret++
			}
		}
	}
	return ret
}

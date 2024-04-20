func bfs(groups *[][]int, visited [][]bool, land, dir [][]int, start []int) {
	ret := []int{start[0], start[1]}
	queue := [][]int{start}
	var curr []int
	visited[start[0]][start[1]] = true
	for len(queue) > 0 {
		curr = queue[0]
		queue = queue[1:]
		for _, d := range dir {
			y := curr[0] + d[0]
			x := curr[1] + d[1]
			if x >= 0 && x < len(land[0]) && y >= 0 && y < len(land) && land[y][x] == 1 && !visited[y][x] {
				queue = append(queue, []int{y, x})
				visited[y][x] = true
			}
		}
	}
	ret = append(ret, curr...)
	*groups = append(*groups, ret)
}

func findFarmland(land [][]int) [][]int {
	ret := [][]int{}
	if len(land) == 0 {
		return ret
	}
	visited := make([][]bool, len(land))
	for i := range visited {
		visited[i] = make([]bool, len(land[0]))
	}
	directions := [][]int{{1, 0}, {0, 1}}
	for y := range land {
		for x := range land[0] {
			if land[y][x] == 1 && !visited[y][x] {
				bfs(&ret, visited, land, directions, []int{y, x})
			}
		}
	}
	return ret
}

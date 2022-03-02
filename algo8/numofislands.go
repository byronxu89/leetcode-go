package algo8

func numIslands(grid [][]byte) int {

	rows := len(grid)
	cols := len(grid[0])
	set := Construct(rows * cols)
	totalcount := (rows * cols)
	direction := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	zerocnt := 0
	var node func(r, c int) int
	node = func(r, c int) int {
		return r*cols + c
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '0' {
				zerocnt++
				continue
			} else {
				for _, elem := range direction {
					nr := i + elem[0]
					nc := j + elem[1]
					if 0 <= nr && nr < rows && 0 <= nc && nc < cols && grid[nr][nc] == '1' {
						flag := set.Join(node(i, j), node(nr, nc))
						if flag {
							totalcount--
						}
					}
				}

			}
		}
	}
	return totalcount - zerocnt

}

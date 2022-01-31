package algo4

func solve(board [][]byte) {
	rowlen := len(board)
	collen := len(board[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= rowlen || y < 0 || y >= collen || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A'
		dfs(x, y+1)
		dfs(x, y-1)
		dfs(x+1, y)
		dfs(x-1, y)
	}

	for i := 0; i < rowlen; i++ {
		dfs(i, 0)
		dfs(i, collen-1)

	}
	for j := 0; j < collen; j++ {
		dfs(0, j)
		dfs(rowlen-1, j)
	}
	for i := 0; i < rowlen; i++ {
		for j := 0; j < collen; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
)

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	R, C := len(board), len(board[0])
	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= R || c < 0 || c >= C || board[r][c] != 'O' {
			return
		}
		board[r][c] = '#'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}
	for i := 0; i < R; i++ {
		dfs(i, 0)
		dfs(i, C-1)
		for j := 0; j < C; j++ {
			dfs(0, j)
			dfs(R-1, j)
		}
	}

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if board[r][c] == 'O' {
				board[r][c] = 'X'
			} else if board[r][c] == '#' {
				board[r][c] = 'O'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}

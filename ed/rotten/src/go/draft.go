package main

import "fmt"

type State struct {
	r, c, mins int
}

func orangesRotting(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var queue []State
	freshCount := 0

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, State{r, c, 0})
			} else if grid[r][c] == 1 {
				freshCount++
			}
		}
	}

	if freshCount == 0 {
		return 0
	}

	minutesPassed := 0
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		minutesPassed = curr.mins

		for _, d := range dirs {
			nr := curr.r + d[0]
			nc := curr.c + d[1]

			if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
				grid[nr][nc] = 2
				freshCount--
				queue = append(queue, State{nr, nc, curr.mins + 1})
			}
		}
	}
	if freshCount == 0 {
		return minutesPassed
	}
	return -1
}

func main() {
	var m, n int

	if _, err := fmt.Scan(&m, &n); err != nil {
		return
	}
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&grid[i][j])
		}
	}
	fmt.Println(orangesRotting(grid))
}

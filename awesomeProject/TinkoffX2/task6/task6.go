package main

import (
	"fmt"
)

type Cell struct {
	x, y int
}

func main() {
	var n int
	fmt.Scan(&n)

	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		var row string
		fmt.Scan(&row)
		for j, char := range row {
			board[i][j] = char
		}
	}

	start, finish := findStartAndFinish(board)

	minMoves := bfs(board, start, finish, n)

	fmt.Println(minMoves)
}

func findStartAndFinish(board [][]rune) (start, finish Cell) {
	for i, row := range board {
		for j, char := range row {
			switch char {
			case 'S':
				start = Cell{i, j}
			case 'F':
				finish = Cell{i, j}
			}
		}
	}
	return start, finish
}

func isValidMove(board [][]rune, cell Cell) bool {
	n := len(board)
	x, y := cell.x, cell.y
	if x < 0 || x >= n || y < 0 || y >= n || board[x][y] == 'G' {
		return false
	}
	return true
}

func bfs(board [][]rune, start, finish Cell, n int) int {
	directions := []Cell{{-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {-2, -1}, {-2, 1}, {2, -1}, {2, 1}}

	queue := []Cell{start}
	visited := make(map[Cell]bool)
	visited[start] = true

	moves := 0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]

			if current == finish {
				return moves
			}

			for _, dir := range directions {
				next := Cell{current.x + dir.x, current.y + dir.y}

				if isValidMove(board, next) && !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		moves++
	}

	return -1 // Если достигли конца и не нашли путь
}

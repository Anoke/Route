package main

import "fmt"

var cell [][]rune

func main() {
	var n int
	fmt.Scan(&n)

	cell = make([][]rune, n)
	for i := 0; i < n; i++ {
		var row string
		fmt.Scan(&row)
		cell[i] = make([]rune, 3)
		for j := 0; j < 3; j++ {
			cell[i][j] = rune(row[j])
		}
	}

	maxMushroom := 0
	for i := 0; i < 3; i++ {
		if cell[0][i] == 'C' {
			maxMushroom = max(maxMushroom, dfs(0, i))
			continue
		}
		if cell[0][i] == '.' {
			maxMushroom = max(maxMushroom, dfs(0, i))
		}
	}

	fmt.Println(maxMushroom)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(x, y int) int {
	mushroom := 0
	if cell[x][y] == 'C' {
		mushroom++
	}

	maxmushroom := 0
	for _, nextY := range []int{y - 1, y, y + 1} {
		if nextY >= 0 && nextY < 3 && x < len(cell)-1 && (cell[x+1][nextY] == 'C' || cell[x+1][nextY] == '.') && cell[x+1][nextY] != 'W' {
			subMushroom := dfs(x+1, nextY)
			maxmushroom = max(maxmushroom, subMushroom)
		}
	}
	return mushroom + maxmushroom
}

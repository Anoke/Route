package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	swappedMatrix := make([][]int, m)
	for i := 0; i < m; i++ {
		swappedMatrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			swappedMatrix[i][j] = matrix[n-j-1][i]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(swappedMatrix[i][j], " ")
		}
		fmt.Println()
	}
}

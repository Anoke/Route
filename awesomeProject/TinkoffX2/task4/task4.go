package main

import "fmt"

func main() {
	var n int
	var direction string
	fmt.Scan(&n, &direction)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	var operations [][]int

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			operations = append(operations, []int{i, j, j, i})
		}
	}

	if direction == "L" {
		for i := 0; i < n/2; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j], matrix[n-i-1][j] = matrix[n-i-1][j], matrix[i][j]
				operations = append(operations, []int{n - i - 1, j, i, j})
			}
		}
	} else {
		for i := 0; i < n; i++ {
			for j := 0; j < n/2; j++ {
				matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
				operations = append(operations, []int{i, n - j - 1, i, j})
			}
		}
	}

	fmt.Println(len(operations))

	for _, op := range operations {
		fmt.Println(op[0], op[1], op[2], op[3])
	}
}

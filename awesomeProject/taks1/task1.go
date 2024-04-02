package main

import (
	"fmt"
)

func checkShips(board []int) bool {
	count1, count2, count3, count4 := 0, 0, 0, 0

	for _, size := range board {
		switch size {
		case 1:
			count1++
		case 2:
			count2++
		case 3:
			count3++
		case 4:
			count4++
		}
	}

	return count1 == 4 && count2 == 3 && count3 == 2 && count4 == 1
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		board := make([]int, 10)
		for j := 0; j < 10; j++ {
			fmt.Scan(&board[j])
		}

		if checkShips(board) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

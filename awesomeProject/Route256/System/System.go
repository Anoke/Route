package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var input string
		fmt.Scan(&input)
		characters := []rune(input)
		result := true

		if len(characters) < 2 || characters[0] != 'M' || characters[len(characters)-1] != 'D' || characters[1] == 'M' {
			result = false
		}

		for j := 1; j < len(characters)-1; j++ {
			switch characters[j] {
			case 'M':
				if characters[j+1] == 'M' {
					result = false
				}
			case 'R':
				if characters[j+1] != 'C' {
					result = false
				}
			case 'C':
				if characters[j+1] != 'M' {
					result = false
				}
			case 'D':
				if characters[j+1] != 'M' {
					result = false
				}
			default:
				result = false
			}
		}

		if result {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		result = true
	}
}

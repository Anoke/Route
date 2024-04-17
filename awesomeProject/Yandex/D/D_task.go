package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var commands string
	fmt.Scan(&commands)

	answer := 0
	for _, command := range commands {
		switch command {
		case 'F':
			answer++
		case 'R':

		case 'L':

		}
	}
}

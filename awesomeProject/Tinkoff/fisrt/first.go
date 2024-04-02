package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func canFormTinkoff(s string) bool {
	requiredLetters := "TINKOFF"
	letterCount := make(map[rune]int)

	for _, char := range s {
		letterCount[char]++
	}

	for _, char := range requiredLetters {
		if count, exists := letterCount[char]; !exists || count == 0 {
			return false
		}
		letterCount[char]--
	}
	for _, i := range letterCount {
		if i != 0 {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		scanner.Scan()
		s := scanner.Text()

		if canFormTinkoff(strings.ToUpper(s)) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

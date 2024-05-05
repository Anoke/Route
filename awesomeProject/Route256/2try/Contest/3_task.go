package main

import (
	"bufio"
	"fmt"
	"os"
)

func IsValidString(s string) bool {
	mostFrequentChar := MostFrequentChar(s)
	stack := []rune{}

	for _, c := range s {
		if len(stack) > 1 && c == stack[len(stack)-2] && c == mostFrequentChar {
			stack = stack[:len(stack)-1]
		} else if len(stack) > 1 && c == stack[len(stack)-1] && c == mostFrequentChar {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	if len(stack) > 1 && stack[len(stack)-1] == stack[len(stack)-2] {
		stack = stack[:len(stack)-2]
	}
	return len(stack) == 0 || len(stack) == 1
}
func MostFrequentChar(s string) rune {
	charCount := make(map[rune]int)

	for _, char := range s {
		charCount[char]++
	}

	var mostFrequentChar rune
	maxCount := 0

	for char, count := range charCount {
		if count > maxCount {
			mostFrequentChar = char
			maxCount = count
		}
	}

	return mostFrequentChar
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var a int
	fmt.Fscanln(reader, &a)

	for i := 0; i < a; i++ {
		var str string
		fmt.Fscanln(reader, &str)

		if IsValidString(str) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
	writer.Flush()
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func Terminal(str string) []string {
	strings := 1
	var result []string
	if len(str) == 0 {
		return result
	}
	result = append(result, "")
	cursor := []int{0, 0}
	for _, char := range str {
		switch {
		case ('a' <= char && char <= 'z') || ('0' <= char && char <= '9'):
			result[cursor[0]] = result[cursor[0]][:cursor[1]] + string(char) + result[cursor[0]][cursor[1]:]
			cursor[1]++
		case char == 'L':
			if cursor[1] > 0 {
				cursor[1]--
			}
		case char == 'R':
			if cursor[1] != len(result[cursor[0]]) && cursor[1] >= 0 {
				cursor[1]++
			}
		case char == 'U':
			if cursor[0] == 0 {
				continue
			}
			if cursor[1] <= len(result[cursor[0]-1]) {
				cursor[0]--
				continue
			}
			if len(result[cursor[0]]) <= len(result[cursor[0]-1]) {
				cursor[0]--
				continue
			}
			cursor[0]--
			cursor[1] = len(result[cursor[0]])
		case char == 'D':
			if cursor[0] == len(result)-1 {
				continue
			}
			if cursor[1] <= len(result[cursor[0]+1]) {
				cursor[0]++
				continue
			}
			if len(result[cursor[0]]) <= len(result[cursor[0]+1]) {
				cursor[0]++
				continue
			}

			cursor[0]++
			cursor[1] = len(result[cursor[0]])

		case char == 'B':
			cursor[1] = 0
		case char == 'E':
			cursor[1] = len(result[cursor[0]])
		case char == 'N':
			strings++
			if cursor[0] < len(result) {
				rightPart := result[cursor[0]][cursor[1]:]
				result[cursor[0]] = result[cursor[0]][:cursor[1]]
				cursor[0]++
				if cursor[0] == len(result) {
					result = append(result, "")
				} else {
					result = append(result[:cursor[0]], append([]string{""}, result[cursor[0]:]...)...)
				}
				result[cursor[0]] = rightPart
				cursor[1] = 0
			}
		}
	}
	return result
}

func main() {
	var iterations int
	fmt.Scan(&iterations)

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 0; i < iterations; i++ {
		text, _ := in.ReadString('\n')
		result := Terminal(text)
		for i := 0; i < len(result); i++ {
			fmt.Fprintln(out, result[i])
		}
		fmt.Fprintln(out, "-")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	t, _ := reader.ReadString('\n')

	words := regexp.MustCompile(`\w+`).FindAllString(t, -1)

	maxLength := 0
	for _, word := range words {
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}
	length := maxLength * 3

	t = regexp.MustCompile(`\s+`).ReplaceAllString(t, " ")

	t = regexp.MustCompile(`,\b`).ReplaceAllString(t, ", ")

	t = strings.Replace(t, " ,", ",", -1)

	t = t + " "

	i := 0
	for i < len(t) {
		if i+length+1 > len(t) {
			length = len(t) - i - 1
		}
		line := t[i : i+length+1]
		if line[len(line)-1] != ' ' {
			lastSpace := strings.LastIndex(line, " ")
			if lastSpace != -1 {
				line = line[:lastSpace+1]
			}
		}
		fmt.Println(strings.TrimSpace(line))
		i += len(line)
	}
}

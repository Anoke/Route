package main

import (
	"fmt"
	"strings"
)

func CheckFirstType(number string) bool {
	if len(number) != 5 {
		return false
	}
	return 'A' <= number[0] && number[0] <= 'Z' && '0' <= number[1] && number[1] <= '9' && '0' <= number[2] && number[2] <= '9' &&
		'A' <= number[3] && number[3] <= 'Z' && 'A' <= number[4] && number[4] <= 'Z'
}

func CheckSecondType(number string) bool {
	if len(number) != 4 {
		return false
	}
	return 'A' <= number[0] && number[0] <= 'Z' && '0' <= number[1] && number[1] <= '9' && 'A' <= number[2] && number[2] <= 'Z' &&
		'A' <= number[3] && number[3] <= 'Z'
}

func FindNumbers(str string) string {
	var result strings.Builder
	for i := 0; i < len(str); {
		if i+5 <= len(str) && CheckFirstType(str[i:i+5]) {
			result.WriteString(str[i : i+5])
			result.WriteRune(' ')
			i += 5
		} else if i+4 <= len(str) && CheckSecondType(str[i:i+4]) {
			result.WriteString(str[i : i+4])
			result.WriteRune(' ')
			i += 4
		} else {
			return "-"
		}
	}
	return strings.TrimSpace(result.String())
}
func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var str string
		fmt.Scan(&str)
		result := FindNumbers(str)
		fmt.Println(result)
	}
}

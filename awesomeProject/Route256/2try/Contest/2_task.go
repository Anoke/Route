package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsValid(letters []string, pin string) string {
	mapa := make(map[string]int)
	for _, letter := range letters {
		mapa[letter]++
	}

	for _, char := range pin {
		letter := string(char)
		if quantity, ok := mapa[letter]; !ok || quantity <= 0 {
			return "NO"
		}
		mapa[letter]--
	}
	for _, count := range mapa {
		if count > 0 {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	firstLine, _ := reader.ReadString('\n')
	firstLine = strings.TrimSpace(firstLine)
	firstLineParts := strings.Fields(firstLine)
	b, _ := strconv.Atoi(firstLineParts[1])

	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strs := strings.Fields(str)
	for i := 0; i < b; i++ {
		pin, _ := reader.ReadString('\n')
		pin = strings.TrimSpace(pin)
		result := IsValid(strs, pin)
		fmt.Println(result)
	}
}

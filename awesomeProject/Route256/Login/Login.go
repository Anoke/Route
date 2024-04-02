package main

import (
	"bufio"
	"fmt"
	"os"
)

func similar(дщ, newLogin string) bool {
	if login1 == login2 {
		return true
	}
	if len(login1) != len(login2) {
		return false
	}

	foundDifference := false
	temp := 0
	for i := 0; i < len(login1)-1; i++ {
		if login1[i] != login2[i] {
			if foundDifference {
				return false
			}

			if login1[i] == login2[i+1] && login1[i+1] == login2[i] {
				foundDifference = true
				temp = i
			} else {
				return false
			}
			i++
		}
	}
	if login1[len(login1)-1] != login2[len(login2)-1] && len(login2)-1 != temp+1 && foundDifference {
		return false
	}

	return foundDifference
}

func main() {
	var n int
	fmt.Scan(&n)

	currentLogins := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&currentLogins[i])
	}

	var m int
	fmt.Scan(&m)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for j := 0; j < m; j++ {
		var newLogin string
		fmt.Scan(&newLogin)

		found := false
		for _, currentLogin := range currentLogins {
			if similar(currentLogin, newLogin) {
				found = true
				break
			}
		}

		if found {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 0)
		}
	}
}

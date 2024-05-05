package main

import (
	"bufio"
	"fmt"
	"os"
)

func similar(oldLogin, newLogin string) bool {

	if len(oldLogin) != len(newLogin) {
		return false
	}

	for i := 0; i < len(newLogin); i++ {
		if newLogin[i] != oldLogin[i] && i != len(newLogin)-1 {
			newLogin = newLogin[:i] + string(newLogin[i+1]) + string(newLogin[i]) + newLogin[i+2:]

			if newLogin == oldLogin {
				return true
			}

			break
		}
	}

	return false
}

func main() {
	var n int
	fmt.Scan(&n)

	currentLogins := make(map[string]bool, n)
	for i := 0; i < n; i++ {
		var login string
		fmt.Scan(&login)
		currentLogins[login] = true
	}

	var m int
	fmt.Scan(&m)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for j := 0; j < m; j++ {
		var newLogin string
		fmt.Scan(&newLogin)

		if currentLogins[newLogin] {
			fmt.Fprintln(out, 1)
		} else {
			found := false
			for currentLogin, _ := range currentLogins {
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
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	fmt.Scan(&str)
	var count int
	fmt.Scan(&count)
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	for i := 0; i < count; i++ {
		defer out.Flush()

		var a, b int
		var s string
		fmt.Fscan(in, &a, &b, &s)
		if a > 0 && a <= len(str) {
			str = str[:a-1] + s + str[b:]
		}
	}
	fmt.Fprint(out, str)
}

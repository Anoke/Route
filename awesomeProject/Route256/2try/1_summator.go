package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	fmt.Fscan(in, &count)
	for i := 0; i < count; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		fmt.Println(a + b)
	}
}

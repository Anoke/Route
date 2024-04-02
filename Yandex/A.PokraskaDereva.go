package main

import "fmt"

func main() {
	var p, v, q, m int
	fmt.Scan(&p, &v)
	fmt.Scan(&q, &m)

	left := 0
	right := 0
	if p-v < q-m {
		left = p - v
	} else {
		left = q - m
	}
	if p+v > q+m {
		right = p + v
	} else {
		right = q + m
	}

	if  {
		
	}
	answer := right - left + 1

	fmt.Println(answer)
}

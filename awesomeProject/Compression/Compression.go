package main

import "fmt"

func Compress(array []int) ([]int, int) {
	var compressed []int
	i := 0
	for i < len(array) {
		start := i
		inc := true

		if i+1 < len(array) && array[i] > array[i+1] {
			inc = false
		}

		for i+1 < len(array) && ((inc && array[i] == array[i+1]-1) || (!inc && array[i] == array[i+1]+1)) {
			i++
		}

		length := i - start
		// Представление отрезка
		if inc {
			compressed = append(compressed, array[start], length)
		} else {
			compressed = append(compressed, array[start], -length)
		}
		i++
	}
	return compressed, len(compressed)
}

func main() {
	var iterations int
	fmt.Scan(&iterations)

	for i := 0; i < iterations; i++ {
		var length int
		fmt.Scan(&length)
		data := make([]int, length)
		for j := 0; j < length; j++ {
			fmt.Scan(&data[j])
		}

		compressed, l := Compress(data)
		fmt.Println(l)
		//fmt.Println(compressed)
		for i := 0; i < len(compressed); i++ {
			fmt.Printf("%d ", compressed[i])
		}
		fmt.Println()
	}
}

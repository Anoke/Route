package main

import "fmt"

func main() {
	var iterations int
	fmt.Scan(&iterations)

	for i := 0; i < iterations; i++ {
		var emp int
		fmt.Scan(&emp)
		minTemp := 15
		maxTemp := 30
		for i := 0; i < emp; i++ {
			var sign string
			var temp int

			fmt.Scan(&sign, &temp)

			if sign == ">=" && temp <= maxTemp && minTemp != -1 {
				if minTemp <= temp {
					minTemp = temp
				}
				fmt.Println(minTemp)
			} else if sign == "<=" && temp >= minTemp && minTemp != -1 {
				if maxTemp >= temp {
					maxTemp = temp
				}
				fmt.Println(minTemp)
			} else {
				minTemp = -1
				fmt.Println(minTemp)
			}
		}
		fmt.Println()
	}
}

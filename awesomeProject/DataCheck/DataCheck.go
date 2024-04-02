package main

import "fmt"

func CheckLeap(year int) bool {
	switch {
	case year%4 == 0 && (year%100 != 0 || year%400 == 0):
		return true
	default:
		return false
	}
}

func IsValidDate(day, month, year int) bool {

	if month <= 12 && month > 0 && year > 1949 && year < 2301 && day > 0 {
		daysInMonth := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		if CheckLeap(year) {
			daysInMonth[2] = 29
		}
		return day <= daysInMonth[month]
	}
	return false
}

func main() {
	var iterations int
	fmt.Scan(&iterations)

	for i := 0; i < iterations; i++ {
		data := make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Scan(&data[j])
		}
		if IsValidDate(data[0], data[1], data[2]) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

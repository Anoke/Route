package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n, p int
		fmt.Scan(&n, &p)

		var totalLost float64
		for j := 0; j < n; j++ {
			var ai int
			fmt.Scan(&ai)

			commission := float64(ai) * float64(p) / 100.0
			lostCents := commission - math.Floor(commission)
			totalLost += lostCents
		}

		result := fmt.Sprintf("%.2f", totalLost)
		fmt.Println(result)
	}
}

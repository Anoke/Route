package main

import "fmt"

func solution(n int, nums []int) int {
	left := 0
	maxFives := 0
	currentFive := 0
	currentDays := 0
	if n < 7 {
		return -1
	}
	for right := 0; right < n; right++ {
		currentDays = right - left + 1
		if currentDays <= 7 {
			if nums[right] == 5 {
				currentFive++
			}
			continue
		}
		if !contains(nums[left:right+1], 2) && !contains(nums[left:right+1], 3) {
			if currentFive > maxFives {
				maxFives = currentFive
			}
		}
		if nums[left] == 5 {
			currentFive--
		}
		left++
		if nums[right] == 5 {
			currentFive++
		}
		right++

	}
	if maxFives == 0 {
		return -1
	}
	return currentFive
}

func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(solution(n, nums))

}

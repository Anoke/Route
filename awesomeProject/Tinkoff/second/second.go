package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "01000"
	common := maximumOddBinaryNumber(s)
	fmt.Println(common)
	//var n, m int
	//fmt.Scan(&n, &m)
	//
	//prices := make([]int, n)
	//for i := 0; i < n; i++ {
	//	fmt.Scan(&prices[i])
	//}
	//
	//// Сначала сортируем цены подарков по возрастанию
	//sort.Ints(prices)
	//
	//// Вычисляем максимальное количество подарков, которое Максим может купить
	//maxGifts := 0
	//for i := 0; i < n; i++ {
	//	if m >= prices[i] {
	//		m -= prices[i]
	//		maxGifts++
	//	} else {
	//		break
	//	}
	//}
	//
	//fmt.Println(m)
}

func maximumOddBinaryNumber(s string) string {
	quantityOne := strings.Count(s, "1")
	indexLast := strings.LastIndex(s, "1")
	if quantityOne == 1 {
		s = s[:indexLast] + s[indexLast+1:] + "1"
		return s
	}
	one := "1"
	zero := "0"
	return strings.Repeat(one, quantityOne-1) + strings.Repeat(zero, len(s)-quantityOne) + "1"
}

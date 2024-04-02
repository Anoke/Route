package main

import (
	"fmt"
	"sort"
	"strings"
)

type PageRange struct {
	Start int
	End   int
}

func findMissingPages(k int, printedPages string) string {
	var missingPages []int

	// Добавляем все страницы, которые необходимо напечатать, в missingPages
	for i := 1; i <= k; i++ {
		missingPages = append(missingPages, i)
	}

	// Удаляем страницы, которые уже были напечатаны
	printed := strings.Split(printedPages, ",")
	for _, page := range printed {
		if strings.Contains(page, "-") {
			// Обработка диапазона страниц
			rangeParts := strings.Split(page, "-")
			start := parsePageNumber(rangeParts[0])
			end := parsePageNumber(rangeParts[1])
			missingPages = removePagesInRange(missingPages, start, end)
		} else {
			// Обработка отдельной страницы
			pageNumber := parsePageNumber(page)
			missingPages = removePage(missingPages, pageNumber)
		}
	}

	// Сортировка страниц, чтобы создать минимальные диапазоны
	sort.Ints(missingPages)

	// Создание диапазонов страниц
	var ranges []PageRange
	for _, page := range missingPages {
		if len(ranges) == 0 || page != ranges[len(ranges)-1].End+1 {
			ranges = append(ranges, PageRange{Start: page, End: page})
		} else {
			ranges[len(ranges)-1].End++
		}
	}

	// Преобразование диапазонов в строку вывода
	var result strings.Builder
	for _, r := range ranges {
		if r.Start == r.End {
			result.WriteString(fmt.Sprintf("%d,", r.Start))
		} else {
			result.WriteString(fmt.Sprintf("%d-%d,", r.Start, r.End))
		}
	}

	// Удаление последней запятой и возврат строки
	return strings.TrimSuffix(result.String(), ",")
}

func parsePageNumber(s string) int {
	// Парсинг строки в число
	var page int
	fmt.Sscanf(s, "%d", &page)
	return page
}

func removePage(pages []int, target int) []int {
	// Удаление страницы из списка
	var result []int
	for _, page := range pages {
		if page != target {
			result = append(result, page)
		}
	}
	return result
}

func removePagesInRange(pages []int, start, end int) []int {
	// Удаление страниц в диапазоне из списка
	var result []int
	for _, page := range pages {
		if page < start || page > end {
			result = append(result, page)
		}
	}
	return result
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var k int
		fmt.Scan(&k)

		var printedPages string
		fmt.Scan(&printedPages)

		result := findMissingPages(k, printedPages)
		fmt.Println(result)
	}
}

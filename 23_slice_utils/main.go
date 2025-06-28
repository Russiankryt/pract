package main

import (
	"fmt"
	"sort"
)

func contains(slice []int, target int) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func filter(slice []int, predicate func(int) bool) []int {
	result := []int{}
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	numbers := []int{5, 2, 8, 1, 7}

	fmt.Println("Исходный срез:", numbers)
	fmt.Println("Содержит 8?", contains(numbers, 8))

	sort.Ints(numbers)
	fmt.Println("Отсортированный:", numbers)

	even := filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Четные числа:", even)
}

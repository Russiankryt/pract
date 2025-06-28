package main

import "fmt"

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func main() {
	var year int
	fmt.Print("Введите год: ")
	fmt.Scanln(&year)

	if isLeapYear(year) {
		fmt.Println(year, "— високосный год")
	} else {
		fmt.Println(year, "— обычный год")
	}
}

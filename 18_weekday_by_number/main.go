package main

import "fmt"

func main() {
	var day int
	fmt.Print("Введите номер дня недели (1–7): ")
	fmt.Scanln(&day)

	switch day {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("Среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница")
	case 6:
		fmt.Println("Суббота")
	case 7:
		fmt.Println("Воскресенье")
	default:
		fmt.Println("Неверный номер дня")
	}
}

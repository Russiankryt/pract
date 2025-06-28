package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Course int
	GPA    float64
}

func printStudent(s Student) {
	fmt.Printf("Имя: %s, Возраст: %d, Курс: %d, Средний балл: %.2f\n", s.Name, s.Age, s.Course, s.GPA)
}

func isExcellent(s Student) bool {
	return s.GPA >= 4.5
}

func main() {
	s := Student{"Андрей", 19, 2, 4.7}
	printStudent(s)

	if isExcellent(s) {
		fmt.Println("Статус: Отличник")
	} else {
		fmt.Println("Статус: Не отличник")
	}
}

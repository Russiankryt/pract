package main

import "fmt"

type Student struct {
	Name string
	Age  int
	GPA  float64
}

func (s Student) Status() string {
	switch {
	case s.GPA >= 4.5:
		return "Отличник"
	case s.GPA >= 3.5:
		return "Хорошист"
	default:
		return "Троечник"
	}
}

func (s Student) IsAdult() bool {
	return s.Age >= 18
}

func main() {
	s := Student{"Димасик", 20, 4.3}
	fmt.Printf("Студент: %s, Возраст: %d, GPA: %.2f\n", s.Name, s.Age, s.GPA)
	fmt.Println("Статус:", s.Status())
	fmt.Println("Совершеннолетний?", s.IsAdult())
}

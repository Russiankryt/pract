package main

import "fmt"

func main() {
	slice := make([]string, 0, 10)

	slice = append(slice, "Первый")
	slice = append(slice, "Второй")
	slice = append(slice, "Третий")
	slice = append(slice, "Четвертый")

	fmt.Println("Элементы среза:")
	for _, elem := range slice {
		fmt.Println(elem)
	}
}

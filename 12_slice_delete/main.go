package main

import "fmt"

func removeElement(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	words := []string{"apple", "banana", "cherry", "date", "elderberry"}
	fmt.Println("Исходный срез:", words)

	words = removeElement(words, 2)
	fmt.Println("После удаления:", words)
}

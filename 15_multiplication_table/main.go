package main

import "fmt"

func main() {
	fmt.Println("Таблица умножения для чисел от 1 до 10:")
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%2d x %2d = %3d   ", i, j, i*j)
		}
		fmt.Println()
	}
}

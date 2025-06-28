package main

import "fmt"

func main() {
	a := 13
	b := 5

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("Сложение: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Вычитание: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Умножение: %d * %d = %d\n", a, b, a*b)
	fmt.Printf("Деление: %d / %d = %d\n", a, b, a/b)
	fmt.Printf("Остаток: %d %% %d = %d\n", a, b, a%b)
}

package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 5
	y := 10

	fmt.Printf("До: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("После: x = %d, y = %d\n", x, y)
}

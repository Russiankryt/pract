package main

import "fmt"

func changeByValue(n int) {
	n = 100
}

func changeByPointer(n *int) {
	*n = 100
}

func main() {
	a := 10
	fmt.Println("До changeByValue:", a)
	changeByValue(a)
	fmt.Println("После changeByValue:", a)

	b := 10
	fmt.Println("До changeByPointer:", b)
	changeByPointer(&b)
	fmt.Println("После changeByPointer:", b)
}

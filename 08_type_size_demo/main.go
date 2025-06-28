package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Размер int:", unsafe.Sizeof(int(0)))
	fmt.Println("Размер float64:", unsafe.Sizeof(float64(0)))
	fmt.Println("Размер bool:", unsafe.Sizeof(true))
	fmt.Println("Размер string:", unsafe.Sizeof("Go"))
	fmt.Println("Размер byte:", unsafe.Sizeof(byte(0)))
	fmt.Println("Размер rune:", unsafe.Sizeof(rune(0)))
}

package main

import "fmt"

func main() {
	var i int = 42
	var f float64 = 3.14
	var b bool = true
	var s string = "GoLang"
	var c complex64 = 1 + 2i
	var r rune = 'G'
	var by byte = 'A'

	fmt.Println("int:", i)
	fmt.Println("float64:", f)
	fmt.Println("bool:", b)
	fmt.Println("string:", s)
	fmt.Println("complex:", c)
	fmt.Println("rune:", r)
	fmt.Println("byte:", by)
}

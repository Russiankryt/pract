package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Привет, GoLang!"
	fmt.Println("Исходная строка:", str)
	fmt.Println("Длина строки:", len(str))
	fmt.Println("Есть 'Go' в строке?:", strings.Contains(str, "Go"))
	fmt.Println("Верхний регистр:", strings.ToUpper(str))
	fmt.Println("Нижний регистр:", strings.ToLower(str))
}

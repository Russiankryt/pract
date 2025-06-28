package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Go — это простой и быстрый язык программирования"
	words := strings.Split(sentence, " ")

	fmt.Println("Слова в предложении:")
	for _, word := range words {
		fmt.Println(word)
	}
}

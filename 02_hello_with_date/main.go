package main

import (
	"fmt"
	"time"
)

func main() {
	name := "Марк"
	currentDate := time.Now().Format("02-01-2006")
	fmt.Printf("Привет, %s! Сегодня %s\n", name, currentDate)
}

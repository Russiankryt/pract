package main

import "fmt"

type Engine struct {
	HorsePower int
	Type       string
}

type Car struct {
	Brand  string
	Model  string
	Engine Engine
}

func main() {
	engine := Engine{HorsePower: 150, Type: "Бензиновый"}
	car := Car{Brand: "Toyota", Model: "Corolla", Engine: engine}

	fmt.Printf("Авто: %s %s\n", car.Brand, car.Model)
	fmt.Printf("Двигатель: %s, %d л.с.\n", car.Engine.Type, car.Engine.HorsePower)
}

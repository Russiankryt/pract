package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	var s Shape

	s = Rectangle{Width: 4, Height: 5}
	fmt.Printf("Площадь прямоугольника: %.2f\n", s.Area())

	s = Circle{Radius: 3}
	fmt.Printf("Площадь круга: %.2f\n", s.Area())
}

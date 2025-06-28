package main

import "fmt"

type Transport interface {
	Drive()
	Stop()
}

type Car struct {
	Brand string
}

func (c Car) Drive() {
	fmt.Println(c.Brand, "поехала.")
}

func (c Car) Stop() {
	fmt.Println(c.Brand, "остановилась.")
}

type Bike struct {
	Model string
}

func (b Bike) Drive() {
	fmt.Println(b.Model, "начал движение.")
}

func (b Bike) Stop() {
	fmt.Println(b.Model, "остановился.")
}

func testTransport(t Transport) {
	t.Drive()
	t.Stop()
}

func main() {
	car := Car{"Lexus"}
	bike := Bike{"best bike"}

	testTransport(car)
	testTransport(bike)
}

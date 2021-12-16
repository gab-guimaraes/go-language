package main

import "fmt"

type Car struct {
	name  string
	brand string
	hp    int
	year  int
}

func toString(car Car) string {
	return car.brand + "\n" + car.name + "\n"
}

func main() {
	lancer := Car{"lancer", "mitsubishi", 370, 2015}
	fmt.Printf(toString(lancer))
}

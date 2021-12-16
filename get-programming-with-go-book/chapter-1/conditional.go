package main

import "fmt"

func main() {
	const PI = 3.14
	var distance = 2000
	fmt.Println(PI)
	fmt.Print(distance)

	var (
		name = "John wick"
	)

	if name == "John wick" {
		fmt.Println("The famous Baba yaga")
	} else {
		fmt.Println("Who is he?")
	}
}

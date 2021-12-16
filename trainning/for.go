package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	treinandoFor()
}

func treinandoFor() {
	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}
}

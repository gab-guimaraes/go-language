package main

import (
	"fmt"
)

func main() {
	cards := []string{"oi", newCard()}
	cards = append(cards, "GTA V")
	fmt.Print(cards) 
}

func newCard() string {
	return "Five of Diamonds"
}

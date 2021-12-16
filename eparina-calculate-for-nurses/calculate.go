package main

import "fmt"

func main() {
	var peso float32 = 70.0
	var bolus float32 = 0.0
	var ttpa float32 = 35.0
	fmt.Println("o peso é", peso, "bolus é", bolus, "o ttpa é", ttpa)
	if ttpa < 35 {
		bolus = 80
	} else if ttpa < 45 {
		bolus = 40
	} else if ttpa < 60 {
		bolus = 40
	} else if ttpa < 85 {
	} else if ttpa < 110 {
	} else {
	}
	var c1 float32 = bolus * peso
	var c2 float32 = c1 / 5000
	concatenated := fmt.Sprint("O resultado é = ", c2, "\n")

	fmt.Print(concatenated)

}

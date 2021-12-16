package main

import (
	"fmt"
)

/*
Short declaration provides an alternative syntax for the var keyword. The following two
lines are equivalent:
var count = 10
count := 10
*/

func main() {
	name := "John Wick"     //implicit type
	var name2 string = "oi" //explicit type
	fmt.Print(name)
	fmt.Print(name2)

}

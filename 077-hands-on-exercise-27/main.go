package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)
	fmt.Printf("The value of x is: %v\n", x)
	fmt.Printf("The value of y is: %v\n", y)

	/*if x < 4 && y < 4 {
		fmt.Printf("x and y are both less than 4\n")
	} else if x > 6 && y > 6 {
		fmt.Printf("x and y are both greater than 6\n")
	} else if x >= 4 && x <= 6 {
		fmt.Printf("x is greater than or equal to 4 and less than or equal to 6\n")
	} else if y != 5 {
		fmt.Printf("y is not 5\n")
	} else {
		fmt.Printf("None of the previous cases were met\n")
	}*/

	switch {
	case x < 4 && y < 4:
		fmt.Printf("x and y are both less than 4\n")
	case x > 6 && y > 6:
		fmt.Printf("x and y are both greater than 6\n")
	case x >= 4 && x <= 6:
		fmt.Printf("x is greater than or equal to 4 and less than or equal to 6\n")
	case y != 5:
		fmt.Printf("y is not 5\n")
	default:
		fmt.Printf("None of the previous cases were met\n")
	}
}

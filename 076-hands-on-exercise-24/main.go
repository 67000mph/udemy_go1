package main

import (
	"fmt"
	"math/rand"
)

// niladic = function without arguments
func init() {
	fmt.Printf("This is where initialization for my program occurs\n")
}

func main() {
	for i := 0; i < 20; i++ {
		randWithCheck(250)
	}
}

func randWithCheck(upperBound int) {
	x := rand.Intn(upperBound + 1)
	fmt.Printf("The value of x is: %v\n", x)
	switch {
	case x <= 100:
		fmt.Printf("between 0 and 100\n")
	case x <= 200:
		fmt.Printf("between 101 and 200\n")
	case x > 200:
		fmt.Printf("between 201 and 250\n")
	}
}

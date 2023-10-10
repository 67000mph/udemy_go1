package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i <= 42; i++ {
		fmt.Printf("Iteration number %v:\n", i)
		x := rand.Intn(5)
		switch {
		case x == 0:
			fmt.Println("x is zero")
		case x == 1:
			fmt.Println("x is one")
		case x == 2:
			fmt.Println("x is two")
		case x == 3:
			fmt.Println("x is three")
		case x == 4:
			fmt.Println("x is four")
		}
	}
}

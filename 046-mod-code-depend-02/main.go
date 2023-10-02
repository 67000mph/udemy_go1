package main

import (
	"fmt"

	"github.com/67000mph/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := puppy.BigBark()
	s4 := puppy.BigBark()
	fmt.Println(s3)
	fmt.Println(s4)
}

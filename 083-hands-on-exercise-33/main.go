package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		if i%2 != 0 {
			fmt.Println(i)
		}
		i++
		time.Sleep(time.Second / 2)
	}
}

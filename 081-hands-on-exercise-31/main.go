package main

import (
	"fmt"
	"time"
)

func main() {
	countdown := 10
	for countdown > 0 {
		fmt.Printf("..%v\n", countdown)
		countdown--
		time.Sleep(time.Second)
	}
	fmt.Printf("BOOM!")
}

package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
		//"Q":          53,
	}

	for k, v := range m {
		fmt.Printf("Key: %v -> Value: %v\n", k, v)
	}

	age1 := m["James"]
	fmt.Println(age1)

	if age2, ok := m["Q"]; ok {
		fmt.Println(age2)
	} else {
		fmt.Println("Not a member of this map")
	}
}

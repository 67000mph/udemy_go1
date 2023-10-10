package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Printf("Key: %v -> Value: %v\n", k, v)
	}

	{
		age := m["James"]
		fmt.Println(age)
	}

	{
		age, ok := m["Q"]
		if ok {
			fmt.Println(age)
		} else {
			fmt.Println("Not a member of this map")
		}
	}
}

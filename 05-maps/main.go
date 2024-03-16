package main

import "fmt"

func main() {

	var names map[string]int

	names = make(map[string]int, 0)

	names["Go"] = 1
	names["Java"] = 2
	names["C++"] = 3

	fmt.Println(names)
	fmt.Println(names["Java"])

	colors := map[string]int{
		"Red":   1,
		"Blue":  2,
		"Black": 3,
	}

	delete(colors, "Blue")
	fmt.Println(colors)
}

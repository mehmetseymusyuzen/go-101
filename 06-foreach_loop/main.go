package main

import (
	"fmt"
)

func main() {

	var numbers = []int{1, 2, 3, 4, 5}

	for index := 0; index < len(numbers); index++ {
		fmt.Println(numbers[index])
	}

	for _, value := range numbers {
		fmt.Println(value)
	}

	language := "Golang"

	for _, character := range language {
		fmt.Printf("%c", character)
	}

	colors := map[string]int{
		"Red":   1,
		"Blue":  2,
		"Black": 3,
	}

	fmt.Println()

	for key, value := range colors {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

}

package main

import "fmt"

func main() {

	// Array
	var names [3]string

	names[0] = "Go"
	names[1] = "Java"
	names[2] = "C++"

	fmt.Println(names)

	var numbers = [4]int{1, 2, 3, 4}
	fmt.Println(numbers)
	numbers[0] = 4
	fmt.Println(numbers[0:3])

	// Slice
	var colors = []string{"Red", "Blue", "Black"}
	fmt.Println(colors)

	colors = append(colors, "Yellow")
	fmt.Println(colors)

}

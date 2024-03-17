package main

import (
	"fmt"
)

func main() {

	total := add(10, 20)
	fmt.Println(total)

	total, difference, multiply := calculation(3, 4)
	fmt.Println(total)
	fmt.Println(difference)
	fmt.Println(multiply)

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))

}

func calculation(x int, y int) (int, int, int) {
	return x + y, x - y, x * y
}

func add(x int, y int) int {
	return x + y
}

func sum(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

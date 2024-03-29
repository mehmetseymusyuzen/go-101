package main

import "fmt"

func main() {

	index := 0

	for index <= 10 {
		index++
	}

	total := 0
	for i := 1; i <= 10; i++ {
		total += i
	}
	fmt.Printf("Total: %d\n", total)

	var names = [3]string{"Go", "Java", "C++"}

	for counter := 0; counter < len(names); {
		fmt.Println(names[counter])
		counter++
	}

}

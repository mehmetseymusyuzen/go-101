package main

import "fmt"

func main() {

	var a int
	a = 10

	var p *int
	p = &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(p)
	*p += 2
	fmt.Println(a)

	x := 2
	fmt.Println("Value of x : ", x)
	passByValue(x)
	fmt.Println("After Pass by Value : ", x)
	passByReference(&x)
	fmt.Println("After Pass by Reference : ", x)

}

func passByValue(x int) {
	x += x

}

func passByReference(x *int) {
	*x += *x
}

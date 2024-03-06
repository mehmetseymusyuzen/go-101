package main

import (
	"fmt"
	"reflect"
)

func main() {

	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Laptop"
	quantity = 21
	discount = 0.50
	isInStock = true

	/*
	    fmt.Println(productName, reflect.TypeOf(productName))
	   	fmt.Println(quantity, reflect.TypeOf(quantity))
	   	fmt.Println(discount, reflect.TypeOf(discount))
	   	fmt.Println(isInStock, reflect.TypeOf(isInStock))
	*/

	/*	fmt.Println("Product name:", productName, "Quantity", quantity, "Discount", discount,
		"isInStock", isInStock)*/

	fmt.Printf("Product name: %s, Quantity: %d, Discount: %f, IsInStock: %t\n", productName, quantity, discount, isInStock)

	var productName2 = "Mouse"

	fmt.Println(productName2, reflect.TypeOf(productName2))

	productNme := "Keyboard"
	fmt.Println(productNme, reflect.TypeOf(productNme))

}

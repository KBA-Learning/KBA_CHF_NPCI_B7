package main

import "fmt"

func main() {
	fmt.Println("Data Types in Go")

	var number1 int
	var number2 float64
	var complex complex64
	var name string
	var IsTrue bool

	fmt.Println(number1, number2, complex, IsTrue, name)
	fmt.Printf("An empty string %q \n", name)

	fmt.Printf("Type of variable is: %T \n", number1)
}
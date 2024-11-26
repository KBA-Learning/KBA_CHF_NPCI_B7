package main

import "fmt"

// Add adds two integers and returns the result.
func Add(a, b int) int {
    return a + b
}

// Subtract subtracts the second integer from the first and returns the result.
func Subtract(a, b int) int {
    return a - b
}

func main() {
	a := Add(5,6)

	b := Subtract(6,2)

	fmt.Printf("Addition %v\n",a)

	fmt.Printf("Subtraction %v\n",b)
}
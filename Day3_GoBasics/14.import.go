package main

import (
	"fmt"
	"basics/greetings" // Import the greetings package
)

func main() {
	// Call the SayHello function from the greetings package
	message := greetings.SayHello("Kumar")
	fmt.Println(message)
}

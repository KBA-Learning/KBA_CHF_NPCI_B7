package main

import "fmt"

func main()  {
	fmt.Println("Introduction to Go")
	fmt.Println("Hello world!")
	var name string
	fmt.Print("Please enter Your Name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello %s good morning \n", name)
	var age = 20
	fmt.Printf("Age of the person is %d years old \n", age)
	email := "kba@gmail.com"
	fmt.Printf("Email Id is %s \n", email)
}
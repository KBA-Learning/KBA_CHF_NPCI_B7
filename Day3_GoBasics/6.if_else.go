package main

import (
	"fmt"
	"log"
)

func main(){
	number := 24
	if number >=0 && number <=10{
		// log.Panicf("Cannot be between 0 to 10")
		fmt.Printf("%d is between 0 and 10 \n", number)
	}else if number >10 && number <= 20{
		fmt.Printf("%d is between 10 and 21 \n", number)
	}else{
		fmt.Printf("Number is not between 0 and 20 \n")
	}
}
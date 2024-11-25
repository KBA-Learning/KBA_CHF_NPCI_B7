package main

import "fmt"

func main() {

	var p *int
	i := 96
	p = &i
	fmt.Println(*p)
	fmt.Println(&p)
	*p = 764
	fmt.Println(i)
}

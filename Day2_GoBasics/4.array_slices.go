package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	var fruits [10]string
	fruits[3] = "Mango"
	fruits[9] = "Apple"
	fmt.Println(fruits, fruits[1])
	numbers := [10]int{1, 4, 7, 5, 8, 9, 0}
	fmt.Println(numbers)

	fmt.Println("Slices")
	var ages []int
	ages = append(ages, 25, 75, 66)
	fmt.Println(ages)
	ages = append(ages, 12,57,87,98)
	fmt.Println(ages)
	names := []string{"Tommy", "Danny", "Robin"}
	fmt.Println(names)
	random := numbers[2:6]
	fmt.Println(random)

	fmt.Println("Length of slices")
	fmt.Println( "length of numbers array: ", len(numbers))
	fmt.Println( "length of slice 'random': ", len(random))
}

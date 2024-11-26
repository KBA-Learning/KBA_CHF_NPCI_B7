package main

import "fmt"

func main(){

	rect1 := Rectangle{3, 5}
	rect2 := Rectangle{12, 4}
	circle1 := Circle{4}
	circle2 := Circle{5}
	totalArea := TotalArea([]Shape{rect1, rect2, circle1, circle2})
	fmt.Println(totalArea)
}


type Shape interface {
	Area() float64
}

type Rectangle struct {
	Breadth float64
	Width   float64
}

func (r Rectangle) Area() float64 {
	return r.Breadth * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func TotalArea(s []Shape) float64 {
	area := 0.0
	for _, v := range s {
		area += v.Area()
	}
	return area
}

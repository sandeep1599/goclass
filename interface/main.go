package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Area() float64 {
	return r.height * r.width
}

func (r *Rectangle) Perimeter() float64 {
	return 2*r.height + 2*r.width
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	r1 := &Rectangle{
		width:  5,
		height: 4,
	}

	PrintAreaAndPerimeter(r1)
	c1 := &Circle{
		radius: 3,
	}

	PrintAreaAndPerimeter(c1)

}

func PrintAreaAndPerimeter(s Shape) {
	defer fmt.Println("Hello world!")
	defer fmt.Println("Hello Dave!")
}

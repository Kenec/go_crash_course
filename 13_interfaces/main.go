package main

import (
	"fmt"
	"math"
)

// Define the area interface
type Shape interface {
	area() float64
}

// Define the derivative interface
type Change interface {
	derivative() float64
}

// Define circle struct
type Circle struct {
	x, y, radius float64
}

// Define rectangle struct
type Rectange struct {
	width, height float64
}

// Define derivative
type Derivative struct {
	x, y float64
}

// Define the circle method
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Define the area method for the rectange
func (r Rectange) area() float64  {
	return r.height * r.width
}

// Define the derivative method
func (d Derivative) derivative() float64 {
	change := (d.x + d.y) / d.x
	return change
}

// getArea method
func getArea(s Shape) float64 {
	return s.area()
}

// getDerivative method
func getDerivative(c Change) float64  {
	return c.derivative()
}

func main() {
	circle := Circle{0, 0, 5}
	rectange := Rectange{4, 7}
	derivative := Derivative{4,2}

	fmt.Println(getArea(circle))
	fmt.Println(getArea(rectange))

	fmt.Println(getDerivative(derivative))


}

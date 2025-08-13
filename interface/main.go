package main

import (
	"fmt"
	"math"
)

type shape interface {
	perimeter() float64
	area() float64
}

type rect struct {
	length, width float64
}

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (r rect) area() float64 {
	return r.length * r.width
}

func (r rect) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func printArea(s interface{}) {
	value, ok := s.(shape)
	if ok {
		fmt.Println("Area:", value.area())
	} else {
		fmt.Println("Not a Shape interface")
	}
}

func main() {
	r := rect{
		length: 10,
		width:  20,
	}
	c := circle{
		radius: 1,
	}
	s := square{
		length: 10,
	}
	printArea(s)
	printArea(r)
	fmt.Printf("Rectangle:\nPerimeter: %.2f\nArea:%.2f\n", r.perimeter(), r.area())
	fmt.Printf("Circle:\nPerimeter: %.2f\nArea:%.2f", c.perimeter(), c.area())
}

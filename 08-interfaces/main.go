package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Name   string
	Radius float64
}

type Rectangle struct {
	Name          string
	Width, Height float64
}

type Shape interface {
	area() float64
	changeName(newName string)
}

func main() {
	circle := Circle{Radius: 4.0}
	rectangle := Rectangle{Width: 3.0, Height: 6.0}

	shapes := []Shape{&circle, &rectangle}

	for _, shape := range shapes {
		fmt.Println("area:", shape.area())
	}
}

func (c *Circle) changeName(newName string) {
	c.Name = newName

}

func (r *Rectangle) changeName(newName string) {
	r.Name = newName
}

func (c *Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r *Rectangle) area() float64 {
	return r.Height * r.Width
}

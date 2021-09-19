package main

import (
	"fmt"
	"math"

	"github.com/sixfwa/crashcourse/randomstuff"
	"github.com/sixfwa/crashcourse/shapes"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(math.Pi)

	circle := shapes.Circle{Radius: 4.5}
	rectangle := shapes.Rectangle{Width: 10.0, Height: 13.0}

	fmt.Println("area of circle:", circle.Area())
	fmt.Println("area of rectangle:", rectangle.Area())
	fmt.Println("is the rectangle square:", rectangle.IsSquare())
	randomstuff.PrintSomeRubbish()
}

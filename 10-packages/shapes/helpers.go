package shapes

import "math"

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) IsSquare() bool {
	return r.Width == r.Height
}

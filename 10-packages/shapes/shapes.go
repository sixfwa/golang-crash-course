package shapes

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

type Shape interface {
	Area() float64
}

package shape

import "math"

// Define the interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle type that implements the Shape interface
type Circle struct {
	Radius float64
}

// Rectangle type that implements the Shape interface
type Rectangle struct {
	Length, Width float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

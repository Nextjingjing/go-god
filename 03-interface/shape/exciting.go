package shape

// Triangle type that implements the Shape interface
type Triangle struct {
	Base, Height float64
}

// Trapezoid type that implements the Shape interface
type Trapezoid struct {
	BaseTop    float64
	BaseBottom float64
	Height     float64
	SideLeft   float64
	SideRight  float64
}

func (t *Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// No Implementation for Perimeter method.
// This will cause a compile-time error if we try to use Triangle as a Shape interface.
// func (t *Triangle) Perimeter() float64 {

func (t *Trapezoid) Area() float64 {
	return 0.5 * (t.BaseTop + t.BaseBottom) * t.Height
}

func (t *Trapezoid) Perimeter() float64 {
	return t.BaseTop + t.BaseBottom + t.SideLeft + t.SideRight
}

// Additional method not related to Shape interface
func (t *Triangle) Test() float64 {
	return 0.0
}

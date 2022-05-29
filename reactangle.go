package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

// CalcPerimeter returns calculation result of perimeter
func (rectangle *Rectangle) CalcPerimeter() float64 {
	return (rectangle.Height + rectangle.Weight) * 2
}

// CalcArea returns calculation result of area
func (rectangle *Rectangle) CalcArea() float64 {
	return rectangle.Height * rectangle.Weight
}

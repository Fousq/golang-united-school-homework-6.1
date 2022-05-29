package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

// CalcPerimeter returns calculation result of perimeter
func (triangle *Triangle) CalcPerimeter() float64 {
	return triangle.Side * 3
}

// CalcArea returns calculation result of area
func (triangle *Triangle) CalcArea() float64 {
	return math.Sqrt(3) / 4 * triangle.Side * triangle.Side
}

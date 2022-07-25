package rect

import "fmt"

type Rect struct {
	Length    float64
	Width     float64
	Area      float64
	Perimeter float64
}

func New(l, w float64) *Rect {
	return &Rect{Length: l, Width: w}
}

// if there is a receiver then it is a method
// receivers are attached to types
func (r *Rect) AreaOf() float64 {
	r.Area = r.Length * r.Width
	return r.Area
}

func (r *Rect) PerimeterOf() float64 {
	r.Perimeter = 2 * (r.Length + r.Width)
	return r.Perimeter
}

func (r *Rect) ShowDetails() {
	fmt.Println("Area:", r.Area)
	fmt.Println("Perimeter", r.Perimeter)
}

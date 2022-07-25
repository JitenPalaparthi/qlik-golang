package main

import (
	"fmt"
	"udt/shape/rect"
)

func main() {
	// var r Rect
	// r.Length = 12.25
	// r.Width = 15.45
	// fmt.Println("Area of Rect:", Area(r))
	r1 := rect.New(12.25, 15.45)
	fmt.Println("Area of Rect:", r1.AreaOf())
	fmt.Println("Perimeter of Rect", r1.PerimeterOf())
	r1.ShowDetails()

	r2 := rect.Rect{Length: 12.25, Width: 15.45}
	fmt.Println("Area of Rect:", r2.AreaOf())
	fmt.Println("Perimeter of Rect", r2.PerimeterOf())
	r2.ShowDetails()

}

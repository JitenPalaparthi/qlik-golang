package main

import (
	"fmt"
	"shapes/shape"
)

type MySquare float32

func (ms MySquare) Area() (float32, error) {
	return float32(ms * ms), nil
}

func (ms MySquare) Perimeter() (float32, error) {
	return float32(4 * ms), nil
}

func main() {
	rect1 := &shape.Rect{L: 100.13, W: 120.45}
	square1 := &shape.Square{S: 10.25}
	rect2 := &shape.Rect{L: 10.13, W: 12.45}
	square2 := &shape.Square{S: 100.25}
	// Area(rect1)
	// Perimeter(rect1)
	// Area(square1)
	// Perimeter(square1)
	var ms MySquare = 12.45
	ss := make([]shape.IShape, 5)
	ss[0] = rect1
	ss[1] = square1
	ss[2] = rect2
	ss[3] = square2
	ss[4] = ms

	// 6th element added here
	ss = append(ss, &shape.Cube{L: 10.23, W: 12.37, H: 9.38})

	for _, v := range ss {
		Area(v)
		Perimeter(v)
	}

}

func Area(iShape shape.IShape) {
	if iShape == nil {
		fmt.Println("nil Interface")
		return
	}
	if area, err := iShape.Area(); err == nil {
		fmt.Println("Area:", area)
	} else {
		fmt.Println(err)
	}

}

func Perimeter(iShape shape.IShape) {
	if iShape == nil {
		fmt.Println("nil Interface")
		return
	}
	if perimeter, err := iShape.Perimeter(); err == nil {
		fmt.Println("Perimeter:", perimeter)
	} else {
		fmt.Println(err)
	}
}

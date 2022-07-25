package main

import (
	"fmt"
	"math/rand"
	r "shapes/shape/rect"
	s "shapes/shape/square"
)

func main() {
	//num := rand.Intn(100)
	fmt.Println("Rect Area:", r.Area(10.25, 14.56))
	fmt.Println("Rect Perimeter:", r.Perimeter(10.25, 14.56))
	fmt.Println("Square Area:", s.Area(25.25))
	fmt.Println("Square Perimeter", s.Perimeter(25.25))
	rand.Intn(100)
}

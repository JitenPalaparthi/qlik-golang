package shape

type IShape interface {
	IArea
	IPerimeter
}

type IArea interface {
	Area() (float32, error)
}

type IPerimeter interface {
	Perimeter() (float32, error)
}

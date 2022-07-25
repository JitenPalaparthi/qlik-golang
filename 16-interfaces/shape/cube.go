package shape

import "errors"

type Cube struct {
	L, W, H float32
}

func (c *Cube) Area() (float32, error) {
	if c == nil {
		return 0, errors.New("nil cube")
	}
	return c.L * c.W * c.H, nil
}

func (c *Cube) Perimeter() (float32, error) {
	if c == nil {
		return 0, errors.New("nil cube")
	}
	return c.L * c.W * c.H, nil
}

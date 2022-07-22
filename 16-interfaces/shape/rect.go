package shape

import "errors"

type Rect struct {
	L, W float32
}

func (r *Rect) Area() (float32, error) {
	if r == nil {
		return 0, errors.New("nil rect")
	}

	return r.L * r.W, nil
}

func (r *Rect) Perimeter() (float32, error) {
	if r == nil {
		return 0, errors.New("nil rect")
	}

	return 2 * (r.L + r.W), nil
}

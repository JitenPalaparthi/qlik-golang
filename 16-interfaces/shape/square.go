package shape

import "errors"

type Square struct {
	S float32
}

func (s *Square) Area() (float32, error) {
	if s == nil {
		return 0, errors.New("nil square")
	}

	return s.S * s.S, nil
}

func (s *Square) Perimeter() (float32, error) {
	if s == nil {
		return 0, errors.New("nil square")
	}

	return 4 * s.S, nil
}

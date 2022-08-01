package models

import "errors"

var (
	ErrInvalidName = errors.New("invalid name field")

	ErrInvalidEmail = errors.New("invalid email field")

	ErrInvalidContactNo = errors.New("invalid contactNo field")
)

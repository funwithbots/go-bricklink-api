package util

import (
	"errors"
)

var (
	ErrNotFound        = errors.New("entity not found")
	ErrForbidden       = errors.New("permission denied")
	ErrNotImplemented  = errors.New("not implemented")
	ErrInvalidArgument = errors.New("invalid argument")
)

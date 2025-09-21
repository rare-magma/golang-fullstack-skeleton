package domain

import "errors"

type InvalidArgument struct {
	error
}

func NewInvalidArgument(msg string) *InvalidArgument {
	return &InvalidArgument{errors.New(msg)}
}

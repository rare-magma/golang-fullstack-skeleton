package domain

import (
	"errors"
	"strings"
)

var ErrEmptyString = errors.New("string value object does not allow empty strings")

type StringValueObject struct {
	ValueObject[string]
}

func NewStringValueObject(value string) (*StringValueObject, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return nil, ErrEmptyString
	}

	vo := NewValueObject[string](trimmed)
	return &StringValueObject{*vo}, nil
}

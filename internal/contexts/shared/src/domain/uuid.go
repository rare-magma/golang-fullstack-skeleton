package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidUUID = errors.New("invalid UUID")

type Uuid struct {
	ValueObject[string]
}

func NewUuid(value string) (*Uuid, error) {
	if _, err := uuid.Parse(value); err != nil {
		return nil, NewInvalidArgument(fmt.Sprintf("<%T> does not allow the value <%s>", Uuid{}, value))
	}

	vo := NewValueObject[string](value)
	return &Uuid{*vo}, nil
}

func RandomUuid() *Uuid {
	uuidStr := uuid.New().String()
	vo, _ := NewUuid(uuidStr)
	return vo
}

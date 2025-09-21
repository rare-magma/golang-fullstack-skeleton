package domain

import (
	"errors"
	"fmt"
	"time"
)

var ErrValueUndefined = errors.New("value must be defined")

type Primitives interface {
	string | int | int32 | int64 | float32 | float64 | bool | time.Time
}

type ValueObject[T Primitives] struct {
	value T
}

func NewValueObject[T Primitives](value T) *ValueObject[T] {
	vo := &ValueObject[T]{value: value}
	vo.ensureValueIsDefined(value)
	return vo
}

func (vo *ValueObject[T]) ensureValueIsDefined(value T) {
	var undefined T
	if value == undefined {
		panic(ErrValueUndefined)
	}
}

func (vo *ValueObject[T]) Value() T {
	return vo.value
}

func (vo *ValueObject[T]) EqualsTo(other ValueObject[T]) bool {
	return other.Value() == vo.value
}

func (vo *ValueObject[T]) String() string {
	return fmt.Sprintf("%v", vo.value)
}
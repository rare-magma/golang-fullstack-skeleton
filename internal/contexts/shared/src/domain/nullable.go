package domain

type Nullable[T any] struct {
	value T
	valid bool
}

func NewNullable[T any](value T) Nullable[T] {
	return Nullable[T]{value: value, valid: true}
}

func Null[T any]() Nullable[T] {
	var zero T
	return Nullable[T]{value: zero, valid: false}
}

func (n Nullable[T]) IsPresent() bool {
	return n.valid
}

func (n Nullable[T]) Get() (T, bool) {
	return n.value, n.valid
}

func (n Nullable[T]) OrElse(fallback T) T {
	if n.valid {
		return n.value
	}
	return fallback
}

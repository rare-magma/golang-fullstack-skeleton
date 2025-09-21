package domain

type Either[L any, R any] struct {
	left  *L
	right *R
}

func Left[L any, R any](value L) Either[L, R] {
	return Either[L, R]{left: &value}
}

func Right[L any, R any](value R) Either[L, R] {
	return Either[L, R]{right: &value}
}

func (e Either[L, R]) IsLeft() bool {
	return e.left != nil
}

func (e Either[L, R]) IsRight() bool {
	return e.right != nil
}

func (e Either[L, R]) LeftValue() (L, bool) {
	if e.left != nil {
		return *e.left, true
	}
	var zero L
	return zero, false
}

func (e Either[L, R]) RightValue() (R, bool) {
	if e.right != nil {
		return *e.right, true
	}
	var zero R
	return zero, false
}

func Fold[L any, R any, T any](e Either[L, R], f func(L) T, g func(R) T) T {
	if e.IsLeft() {
		return f(*e.left)
	}
	return g(*e.right)
}

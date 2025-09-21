package domain

type NumberValueObject struct {
	ValueObject[int]
}

func NewNumberValueObject(value int) (*NumberValueObject, error) {
	vo := NewValueObject(value)
	return &NumberValueObject{*vo}, nil
}

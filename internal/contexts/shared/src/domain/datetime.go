package domain

import (
	"time"
)

type Datetime struct {
	StringValueObject
}

func NewDatetime(datetimeStr string) (*Datetime, error) {
	_, err := time.Parse(time.RFC3339, datetimeStr)
	if err != nil {
		return nil, NewInvalidArgument("Datetime does not allow the value <" + datetimeStr + ">")
	}

	vo, err := NewStringValueObject(datetimeStr)
	if err != nil {
		return nil, err
	}
	return &Datetime{*vo}, nil
}

func (dt *Datetime) Add(seconds int64) *Datetime {
	t, _ := time.Parse(time.RFC3339, dt.Value())
	newTime := t.Add(time.Duration(seconds) * time.Second)
	newDt, _ := NewDatetime(newTime.Format(time.RFC3339))
	return newDt
}
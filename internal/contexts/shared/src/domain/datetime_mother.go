package domain

import (
	"time"

	"github.com/jaswdr/faker/v2"
)

type DatetimeMother struct {
	f *faker.Faker
}

func (dm *DatetimeMother) Create(datetimeStr string) (*Datetime, error) {
	return NewDatetime(datetimeStr)
}

func (dm *DatetimeMother) Random() *Datetime {
	date := dm.f.Time()
	dt, _ := NewDatetime(date.RFC3339(date.Time(time.Now())))
	return dt
}

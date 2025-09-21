package infrastructure

import (
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain"
	"time"
)

type CurrentTimeClock struct{}

func NewCurrentTimeClock() *CurrentTimeClock {
	return &CurrentTimeClock{}
}

func (c *CurrentTimeClock) Now() domain.Datetime {
	time, _ := domain.NewDatetime(time.Now().String())
	return *time
}
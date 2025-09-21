package domain

import (
	"strings"
	"time"

	"github.com/jaswdr/faker/v2"
)

type ObjectMother struct {
	f *faker.Faker
}

func NewObjectMother() *ObjectMother {
	fakerInstance := faker.New()
	return &ObjectMother{
		f: &fakerInstance,
	}
}

func (om *ObjectMother) Repeat(callable func() interface{}, iterations ...int) []interface{} {
	var iter int
	if len(iterations) > 0 {
		iter = iterations[0]
	} else {
		iter = om.ZeroOrPositiveNumber(20)
	}
	results := make([]interface{}, iter)
	for i := 0; i < iter; i++ {
		results[i] = callable()
	}
	return results
}

func (om *ObjectMother) Uuid() string {
	return om.f.UUID().V4()
}

func (om *ObjectMother) IndexNumber(max int) int {
	return om.f.IntBetween(0, max)
}

func (om *ObjectMother) Word() string {
	return om.f.Lorem().Word()
}

func (om *ObjectMother) Words() string {
	return strings.Join(om.f.Lorem().Words(om.f.RandomDigit()), " ")
}

func (om *ObjectMother) Coin() bool {
	return om.f.Bool()
}

func (om *ObjectMother) PositiveNumber(max ...int) int {
	var m int
	if len(max) > 0 {
		m = max[0]
	} else {
		m = 9999
	}
	return om.f.IntBetween(1, m)
}

func (om *ObjectMother) ZeroOrPositiveNumber(max ...int) int {
	var m int
	if len(max) > 0 {
		m = max[0]
	} else {
		m = 9999
	}
	return om.f.IntBetween(0, m)
}

func (om *ObjectMother) Email() string {
	return om.f.Internet().Email()
}

func (om *ObjectMother) PhoneNumber() string {
	return om.f.Phone().Number()
}

func (om *ObjectMother) Text() string {
	return om.f.Lorem().Paragraph(om.f.RandomDigit())
}

func (om *ObjectMother) RecentDate() time.Time {
	return om.f.Time().Time(time.Now().Add(-24 * time.Hour))
}

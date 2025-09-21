package domain

import (
	"time"
)

type DomainEvent interface {
	EventID() string

	AggregateID() string

	OccurredOn() time.Time

	EventName() string

	ToPrimitives() any
}

type DomainEventBase struct {
	eventID     string
	aggregateID string
	occurredOn  time.Time
	eventName   string
}

func NewDomainEventBase(eventName, aggregateID, eventID string, occurredOn time.Time) DomainEventBase {
	if eventID == "" {
		eventID = RandomUuid().Value()
	}
	if occurredOn.IsZero() {
		occurredOn = time.Now().UTC()
	}

	return DomainEventBase{
		eventID:     eventID,
		aggregateID: aggregateID,
		occurredOn:  occurredOn,
		eventName:   eventName,
	}
}

func (d DomainEventBase) EventID() string { return d.eventID }

func (d DomainEventBase) AggregateID() string { return d.aggregateID }

func (d DomainEventBase) OccurredOn() time.Time { return d.occurredOn }

func (d DomainEventBase) EventName() string { return d.eventName }

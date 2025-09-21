package domain

type AggregateRoot struct {
	domainEvents []DomainEvent
}

func NewAggregateRoot() *AggregateRoot {
	return &AggregateRoot{
		domainEvents: make([]DomainEvent, 0),
	}
}

func (ar *AggregateRoot) PullDomainEvents() []DomainEvent {
	events := make([]DomainEvent, len(ar.domainEvents))
	copy(events, ar.domainEvents)
	ar.domainEvents = make([]DomainEvent, 0)
	return events
}

func (ar *AggregateRoot) Record(event DomainEvent) {
	ar.domainEvents = append(ar.domainEvents, event)
}

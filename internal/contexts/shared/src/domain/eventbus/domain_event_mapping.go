package eventbus

import (
	"context"
	"fmt"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain"
)

type DomainEventConstructor interface {
	domain.DomainEvent
	DomainEventName
}

type DomainEventMapping struct {
	mapping map[string]DomainEventConstructor
}

func NewDomainEventMapping(subscribers []DomainEventSubscriber) *DomainEventMapping {
	return &DomainEventMapping{
		mapping: formatSubscribers(subscribers),
	}
}

func formatSubscribers(subscribers []DomainEventSubscriber) map[string]DomainEventConstructor {
	m := make(map[string]DomainEventConstructor)

	for _, subscriber := range subscribers {
		for _, domainEventCtor := range subscriber.SubscribedTo() {
			eventName := domainEventCtor.EventName()
			m[eventName] = domainEventCtor.(DomainEventConstructor)
		}
	}

	return m
}

func (m *DomainEventMapping) For(ctx context.Context, name string) (DomainEventConstructor, error) {
	domainEventCtor, exists := m.mapping[name]
	if !exists {
		return nil, fmt.Errorf("the Domain Event constructor for %s doesn't exist or has no subscribers", name)
	}

	return domainEventCtor, nil
}
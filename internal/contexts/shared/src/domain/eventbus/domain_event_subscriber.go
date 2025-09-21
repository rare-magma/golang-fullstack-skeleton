package eventbus

import (
	"context"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain"
)

type DomainEventSubscriber interface {
	Name() string
	SubscribedTo() []DomainEventName
	On(ctx context.Context, domainEvent domain.DomainEvent) error
}

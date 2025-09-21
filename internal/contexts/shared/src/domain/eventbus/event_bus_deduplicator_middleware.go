package eventbus

import (
	"context"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain"
)

type EventBusDeduplicatorMiddleware struct {
	deduplicator DomainEventDeduplicator
}

func NewEventBusDeduplicatorMiddleware(deduplicator DomainEventDeduplicator) *EventBusDeduplicatorMiddleware {
	return &EventBusDeduplicatorMiddleware{
		deduplicator: deduplicator,
	}
}

func (m *EventBusDeduplicatorMiddleware) Run(ctx context.Context, events []domain.DomainEvent) ([]domain.DomainEvent, error) {
	return m.deduplicator.Deduplicate(ctx, events)
}

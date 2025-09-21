package eventbus

import (
	"context"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain"
)

type DomainEventDeduplicator interface {
	Deduplicate(ctx context.Context, domainEvents []domain.DomainEvent) ([]domain.DomainEvent, error)
	Purge(ctx context.Context, domainEvent domain.DomainEvent) error
}

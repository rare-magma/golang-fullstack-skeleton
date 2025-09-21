package eventbus

import (
	"context"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain"
)

type DomainEventMarshaller interface {
	Marshall(ctx context.Context, event domain.DomainEvent) (any, error)
}
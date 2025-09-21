package eventbus

import (
	"context"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain"
)

type EventBusMiddleware interface {
	Run(ctx context.Context, events []domain.DomainEvent) ([]domain.DomainEvent, error)
}

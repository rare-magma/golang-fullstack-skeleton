package eventbus

import (
	"context"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain"
)

type DomainEventRepository interface {
	Save(ctx context.Context, events ...domain.DomainEvent) error
}

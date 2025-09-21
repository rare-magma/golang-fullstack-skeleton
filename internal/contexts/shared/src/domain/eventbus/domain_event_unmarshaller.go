package eventbus

import (
	"context"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain"
)

type DomainEventUnmarshaller interface {
	Unmarshall(ctx context.Context, infraEvent any) (domain.DomainEvent, error)
}

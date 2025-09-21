package eventbus

import "golang-fullstack-skeleton/internal/contexts/shared/src/domain"

type EventBus interface {
	Publish(events []domain.DomainEvent) error
}

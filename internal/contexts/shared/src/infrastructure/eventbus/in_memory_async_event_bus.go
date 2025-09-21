package eventbus

import (
	"context"
	"fmt"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain/eventbus"
	"sync"
)

type InMemoryAsyncEventBus struct {
	subscribers map[string][]eventbus.DomainEventSubscriber
	mu          sync.RWMutex
}

func NewInMemoryAsyncEventBus() *InMemoryAsyncEventBus {
	return &InMemoryAsyncEventBus{
		subscribers: make(map[string][]eventbus.DomainEventSubscriber),
	}
}

func (b *InMemoryAsyncEventBus) RegisterSubscriber(subscriber eventbus.DomainEventSubscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, eventName := range subscriber.SubscribedTo() {
		b.subscribers[eventName.EventName()] = append(b.subscribers[eventName.EventName()], subscriber)
	}
}

func (b *InMemoryAsyncEventBus) Publish(ctx context.Context, events []domain.DomainEvent) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	var wg sync.WaitGroup
	errChan := make(chan error, len(events)*10)
	var errs []error

	for _, event := range events {
		subscribers, exists := b.subscribers[event.EventName()]
		if !exists {
			errs = append(errs, fmt.Errorf("no subscribers found for event %s", event.EventName()))
			continue
		}

		for _, subscriber := range subscribers {
			wg.Add(1)
			go func(s eventbus.DomainEventSubscriber, e domain.DomainEvent) {
				defer wg.Done()
				if err := s.On(ctx, e); err != nil {
					errChan <- fmt.Errorf("handler failed for event %s: %w", e.EventName(), err)
				}
			}(subscriber, event)
		}
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("encountered %d errors during event publishing: %v", len(errs), errs)
	}

	return nil
}

package eventbus

import (
	"context"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain/eventbus"
	"sync"
	"testing"
	"time"
)

type ExampleEvent struct {
	eventName   string
	aggregateID string
	eventID     string
	occurredOn  time.Time
}

func NewExampleEvent(id string, eventID string, occurredOn time.Time) *ExampleEvent {
	if eventID == "" {
		eventID = domain.RandomUuid().Value()
	}
	if occurredOn.IsZero() {
		occurredOn = time.Now()
	}
	return &ExampleEvent{
		eventName:   "example:event",
		aggregateID: id,
		eventID:     eventID,
		occurredOn:  occurredOn,
	}
}

func (e *ExampleEvent) EventName() string {
	return e.eventName
}

func (e *ExampleEvent) AggregateID() string {
	return e.aggregateID
}

func (e *ExampleEvent) EventID() string {
	return e.eventID
}

func (e *ExampleEvent) OccurredOn() time.Time {
	return e.occurredOn
}

func (e *ExampleEvent) ToPrimitives() any {
	return e
}

type DomainEventSubscriberExample struct {
	name        string
	expectation func(*ExampleEvent)
	wg          sync.WaitGroup
}

func NewDomainEventSubscriberExample() *DomainEventSubscriberExample {
	return &DomainEventSubscriberExample{
		name: "example_subscriber",
	}
}

func (s *DomainEventSubscriberExample) Name() string {
	return s.name
}

func (s *DomainEventSubscriberExample) SubscribedTo() []eventbus.DomainEventName {
	return []eventbus.DomainEventName{&ExampleEvent{eventName: "example:event"}}
}

func (s *DomainEventSubscriberExample) On(ctx context.Context, domainEvent domain.DomainEvent) error {
	s.wg.Add(1)
	defer s.wg.Done()

	if s.expectation != nil {
		s.expectation(domainEvent.(*ExampleEvent))
	}
	return nil
}

func (s *DomainEventSubscriberExample) SetExpectation(fn func(*ExampleEvent)) {
	s.expectation = fn
}

func (s *DomainEventSubscriberExample) Wait() {
	s.wg.Wait()
}

func TestInMemoryAsyncEventBus(t *testing.T) {
	ctx := context.Background()
	objectMother := domain.NewObjectMother()

	t.Run("no subscribers registered returns error", func(t *testing.T) {
		event := NewExampleEvent(objectMother.Uuid(), "", time.Time{})
		eventBus := NewInMemoryAsyncEventBus()

		err := eventBus.Publish(ctx, []domain.DomainEvent{event})
		if err == nil {
			t.Errorf("Expected error when publishing with no subscribers, got: %v", err)
		}
	})

	t.Run("subscriber should be called for subscribed event", func(t *testing.T) {
		event := NewExampleEvent(objectMother.Uuid(), "", time.Time{})
		subscriber := NewDomainEventSubscriberExample()
		eventBus := NewInMemoryAsyncEventBus()

		eventBus.RegisterSubscriber(subscriber)

		called := false
		subscriber.SetExpectation(func(actual *ExampleEvent) {
			called = true
			if actual != event {
				t.Errorf("Expected subscriber to receive event %v, got %v", event, actual)
			}
		})

		err := eventBus.Publish(ctx, []domain.DomainEvent{event})
		if err != nil {
			t.Errorf("Expected no error when publishing, got: %v", err)
		}

		subscriber.Wait()

		if !called {
			t.Error("Expected subscriber to be called, but it was not")
		}
	})
}
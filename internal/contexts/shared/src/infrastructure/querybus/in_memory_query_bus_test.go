package querybus

import (
	"context"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain/querybus"
	"testing"
)

type UnhandledQuery struct{}

func (q UnhandledQuery) QueryName() string {
	return "unhandled.query"
}

type HandledQuery struct{}

func (q HandledQuery) QueryName() string {
	return "handled.query"
}

type MyQueryHandler struct{}

func (h *MyQueryHandler) SubscribedTo() querybus.Query {
	return HandledQuery{}
}

type MyResponse struct {
	data any
}

func (h *MyQueryHandler) Handle(ctx context.Context, query querybus.Query) (querybus.Response, error) {
	return MyResponse{data: "data"}, nil
}

func TestInMemoryQueryBus(t *testing.T) {
	t.Run("throws an error when no handlers are registered", func(t *testing.T) {
		unhandledQuery := UnhandledQuery{}
		queryBus := NewInMemoryQueryBus()
		ctx := context.Background()

		_, err := queryBus.Ask(ctx, unhandledQuery)
		if err == nil || err.Error() != "no query handlers registered" {
			t.Errorf("expected error 'no query handlers registered', got %v", err)
		}
	})

	t.Run("throws an error if dispatches a query without handler", func(t *testing.T) {
		unhandledQuery := UnhandledQuery{}
		queryHandlersInformation := NewQueryHandlersInformation([]querybus.QueryHandler{})
		queryBus := NewInMemoryQueryBus()
		queryBus.RegisterHandlers(queryHandlersInformation)
		ctx := context.Background()

		_, err := queryBus.Ask(ctx, unhandledQuery)
		if err == nil {
			t.Error("expected QueryNotRegisteredError, got nil")
		}
		if _, ok := err.(querybus.QueryNotRegisteredError); !ok {
			t.Errorf("expected QueryNotRegisteredError, got %v", err)
		}
	})

	t.Run("accepts a query with handler", func(t *testing.T) {
		handledQuery := HandledQuery{}
		myQueryHandler := &MyQueryHandler{}
		queryHandlersInformation := NewQueryHandlersInformation([]querybus.QueryHandler{myQueryHandler})
		queryBus := NewInMemoryQueryBus()
		queryBus.RegisterHandlers(queryHandlersInformation)
		ctx := context.Background()

		r, err := queryBus.Ask(ctx, handledQuery)
		if err != nil || r == nil {
			t.Errorf("expected no error, got %v %s", err, r)
		}
	})
}

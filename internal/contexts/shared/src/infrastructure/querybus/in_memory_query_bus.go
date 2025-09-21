package querybus

import (
	"context"
	"errors"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain/querybus"
)

type InMemoryQueryBus struct {
	queryHandlersInformation *QueryHandlersInformation
}

func NewInMemoryQueryBus() *InMemoryQueryBus {
	return &InMemoryQueryBus{}
}

func (b *InMemoryQueryBus) RegisterHandlers(queryHandlersInformation *QueryHandlersInformation) {
	b.queryHandlersInformation = queryHandlersInformation
}

func (b *InMemoryQueryBus) Ask(ctx context.Context, query querybus.Query) (querybus.Response, error) {
	if b.queryHandlersInformation == nil {
		return nil, errors.New("no query handlers registered")
	}

	handler, err := b.queryHandlersInformation.Search(query)
	if err != nil {
		return nil, err
	}
	response, err := handler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	return response, nil
}

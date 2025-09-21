package commandbus

import (
	"context"
	"errors"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain/commandbus"
)

type InMemoryCommandBus struct {
	commandHandlersInformation *CommandHandlersInformation
}

func NewInMemoryCommandBus() *InMemoryCommandBus {
	return &InMemoryCommandBus{}
}

func (b *InMemoryCommandBus) RegisterHandlers(commandHandlersInformation *CommandHandlersInformation) {
	b.commandHandlersInformation = commandHandlersInformation
}

func (b *InMemoryCommandBus) Dispatch(ctx context.Context, command commandbus.Command) error {
	if b.commandHandlersInformation == nil {
		return errors.New("no command handlers registered")
	}

	handler, err := b.commandHandlersInformation.Search(command)
	if err != nil {
		return err
	}

	return handler.Handle(ctx, command)
}

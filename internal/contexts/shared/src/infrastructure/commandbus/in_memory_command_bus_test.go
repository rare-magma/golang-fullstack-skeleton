package commandbus

import (
	"context"
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain/commandbus"
	"testing"
)

type UnhandledCommand struct{}

func (c UnhandledCommand) CommandName() string {
	return "unhandled.command"
}

func (c UnhandledCommand) IsCommand() bool {
	return true
}

type HandledCommand struct{}

func (c HandledCommand) CommandName() string {
	return "handled.command"
}

func (c HandledCommand) IsCommand() bool {
	return true
}

type MyCommandHandler struct{}

func (h MyCommandHandler) SubscribedTo() commandbus.Command {
	return HandledCommand{}
}

func (h MyCommandHandler) Handle(ctx  context.Context, command commandbus.Command) error {
	return nil
}

func TestInMemoryCommandBus(t *testing.T) {
	t.Run("returns an error when no handlers are registered", func(t *testing.T) {
		commandBus := NewInMemoryCommandBus()
		unhandledCommand := UnhandledCommand{}
		ctx := context.Background()

		err := commandBus.Dispatch(ctx, unhandledCommand)

		if err == nil || err.Error() != "no command handlers registered" {
			t.Errorf("Expected error 'no command handlers registered', got %v", err)
		}
	})

	t.Run(" an error if it dispatches a command without handler", func(t *testing.T) {
		commandBus := NewInMemoryCommandBus()
		commandHandlersInformation := NewCommandHandlersInformation([]commandbus.CommandHandler{})
		commandBus.RegisterHandlers(commandHandlersInformation)
		unhandledCommand := UnhandledCommand{}
		ctx := context.Background()

		err := commandBus.Dispatch(ctx, unhandledCommand)
		if err != err.(commandbus.CommandNotRegisteredError) {
				t.Errorf("Expected %v, got %v", err.(commandbus.CommandNotRegisteredError), err)
		}
	})

	t.Run("accepts a command with handler", func(t *testing.T) {
		commandBus := NewInMemoryCommandBus()
		myCommandHandler := MyCommandHandler{}
		commandHandlersInformation := NewCommandHandlersInformation([]commandbus.CommandHandler{myCommandHandler})
		commandBus.RegisterHandlers(commandHandlersInformation)
		handledCommand := HandledCommand{}
		ctx := context.Background()

		err := commandBus.Dispatch(ctx, handledCommand)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}

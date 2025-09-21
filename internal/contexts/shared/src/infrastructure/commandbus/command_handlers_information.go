package commandbus

import (
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain/commandbus"
)

type CommandHandlersInformation struct {
	commandHandlersMap map[string]commandbus.CommandHandler
}

func NewCommandHandlersInformation(commandHandlers []commandbus.CommandHandler) *CommandHandlersInformation {
	chi := &CommandHandlersInformation{
		commandHandlersMap: make(map[string]commandbus.CommandHandler),
	}
	chi.formatHandlers(commandHandlers)
	return chi
}

func (chi *CommandHandlersInformation) formatHandlers(commandHandlers []commandbus.CommandHandler) {
	for _, handler := range commandHandlers {
		chi.commandHandlersMap[handler.SubscribedTo().CommandName()] = handler
	}
}

func (chi *CommandHandlersInformation) Search(command commandbus.Command) (commandbus.CommandHandler, error) {
	handler, exists := chi.commandHandlersMap[command.CommandName()]
	if handler == nil || !exists {
		return nil, commandbus.CommandNotRegisteredError{Command: command}
	}
	return handler, nil
}

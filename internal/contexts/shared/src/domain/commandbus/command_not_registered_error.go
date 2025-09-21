package commandbus

import "fmt"

type CommandNotRegisteredError struct {
	Command Command
}

func (e CommandNotRegisteredError) Error() string {
	return fmt.Sprintf("no handler registered for command: %s", e.Command.CommandName())
}

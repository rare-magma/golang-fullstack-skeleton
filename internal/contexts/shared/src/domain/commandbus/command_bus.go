package commandbus

import "context"

type CommandBus interface {
	Dispatch(ctx context.Context, cmd Command) error
}
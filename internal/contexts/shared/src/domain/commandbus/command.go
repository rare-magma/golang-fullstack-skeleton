package commandbus

type Command interface {
	CommandName() string
	IsCommand() bool
}
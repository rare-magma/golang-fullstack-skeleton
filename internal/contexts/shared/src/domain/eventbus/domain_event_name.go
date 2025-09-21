package eventbus

type DomainEventName interface {
	EventName() string
}
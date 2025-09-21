package arranger

type Arranger interface {
	Arrange() error
	Close() error
}
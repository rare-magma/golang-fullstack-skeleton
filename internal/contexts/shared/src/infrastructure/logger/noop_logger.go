package logger

type NoopLogger struct{}

func NewNoopLogger() *NoopLogger {
	return &NoopLogger{}
}

func (l *NoopLogger) Info(msg string, args ...any) {}
func (l *NoopLogger) Error(msg string, args ...any) {}
func (l *NoopLogger) Debug(msg string, args ...any) {}
func (l *NoopLogger) Warn(msg string, args ...any) {}
package logger

import (
	"log"
	"os"
)

type StandardLogger struct {
	info  *log.Logger
	error *log.Logger
	debug *log.Logger
	warn  *log.Logger
}

func NewStandardLogger() *StandardLogger {
	return &StandardLogger{
		info:  log.New(os.Stdout, "INFO  ", log.LstdFlags|log.Lmsgprefix),
		error: log.New(os.Stderr, "ERROR ", log.LstdFlags|log.Lmsgprefix),
		debug: log.New(os.Stdout, "DEBUG ", log.LstdFlags|log.Lmsgprefix),
		warn:  log.New(os.Stdout, "WARN  ", log.LstdFlags|log.Lmsgprefix),
	}
}

func (l *StandardLogger) Info(msg string, args ...any) {
	l.info.Printf(msg, args...)
}

func (l *StandardLogger) Error(msg string, args ...any) {
	l.error.Printf(msg, args...)
}

func (l *StandardLogger) Debug(msg string, args ...any) {
	l.debug.Printf(msg, args...)
}

func (l *StandardLogger) Warn(msg string, args ...any) {
	l.warn.Printf(msg, args...)
}

package servo

import (
	"fmt"
	"log"
)

// Logger is the interface for logging.
// Disclaimer: This interface was copied from github.com/mgutz/logxi
type Logger interface {
	Trace(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{}) error
	Error(msg string, args ...interface{}) error
	Fatal(msg string, args ...interface{})
	Log(level int, msg string, args []interface{})

	SetLevel(int)
	IsTrace() bool
	IsDebug() bool
	IsInfo() bool
	IsWarn() bool
	// Error, Fatal not needed, those SHOULD always be logged
}

// LoggerProvider is the interface to retrieve a new Logger.
type LoggerProvider interface {

	// GetLogger should return a logger by name.
	// The returned logger may have been cached
	GetLogger(name string) Logger
}

// NullLogger is a null implementation of the Logger interface
type NullLogger struct {}

func (l *NullLogger) Trace(msg string, args ...interface{}) {}
func (l *NullLogger) Debug(msg string, args ...interface{}) {}
func (l *NullLogger) Info(msg string, args ...interface{}) {}
func (l *NullLogger) Warn(msg string, args ...interface{}) error {return fmt.Errorf(msg, args...)}
func (l *NullLogger) Error(msg string, args ...interface{}) error {return fmt.Errorf(msg, args...)}
func (l *NullLogger) Fatal(msg string, args ...interface{}) { log.Fatalf(msg, args...)}
func (l *NullLogger) Log(level int, msg string, args []interface{}) {}
func (l *NullLogger) SetLevel(int) {}
func (l *NullLogger) IsTrace() bool { return false }
func (l *NullLogger) IsDebug() bool { return false }
func (l *NullLogger) IsInfo() bool { return false }
func (l *NullLogger) IsWarn() bool { return false }

// NullLoggerProvider is a null implementation of the LoggerProvider interface
type NullLoggerProvider struct {}

func (p *NullLoggerProvider) GetLogger(_ string) Logger {
	return new(NullLogger)
}

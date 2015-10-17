package servo

import (
	"fmt"
	"io"
	"time"
)

// The log level constants are a copied subset of the log constants of github.com/mgutz/logxi.
const (
	// LevelError is a non-urgen failure to notify devlopers or admins
	LevelError = 3

	// LevelWarn indiates an error will occur if action is not taken, eg file system 85% full
	LevelWarn = 4

	// LevelInfo is info level
	LevelInfo = 6

	// LevelDebug is debug level
	LevelDebug = 7

	// LevelTrace is trace level and displays file and line in terminal
	LevelTrace = 10

	// LevelAll is all levels
	LevelAll = 1000
)

// Logger is the interface for logging.
type Logger interface {
	Trace(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{}) error
	Error(msg string, args ...interface{}) error

	Level() int
	SetLevel(int)
	IsTrace() bool
	IsDebug() bool
	IsInfo() bool
	IsWarn() bool
	IsError() bool
}

// LoggerProvider is the interface to retrieve a new Logger.
type LoggerProvider interface {

	// GetLogger should return a logger by name.
	// The returned logger may have been cached
	GetLogger(name string) Logger
}

// NullLogger is a null implementation of the Logger interface
type NullLogger struct{}

func (l *NullLogger) Trace(msg string, args ...interface{})       {}
func (l *NullLogger) Debug(msg string, args ...interface{})       {}
func (l *NullLogger) Info(msg string, args ...interface{})        {}
func (l *NullLogger) Warn(msg string, args ...interface{}) error  { return fmt.Errorf(msg, args...) }
func (l *NullLogger) Error(msg string, args ...interface{}) error { return fmt.Errorf(msg, args...) }
func (l *NullLogger) SetLevel(int)                                {}
func (l *NullLogger) Level() int                                  { return -1000000 }
func (l *NullLogger) IsTrace() bool                               { return false }
func (l *NullLogger) IsDebug() bool                               { return false }
func (l *NullLogger) IsInfo() bool                                { return false }
func (l *NullLogger) IsWarn() bool                                { return false }
func (l *NullLogger) IsError() bool                               { return false }

// NullLoggerProvider is a null implementation of the LoggerProvider interface
type NullLoggerProvider struct{}

func (p *NullLoggerProvider) GetLogger(_ string) Logger {
	return new(NullLogger)
}

// SimpleLogger is simple implementation of the Logger interface that just prints to stdout.
type SimpleLogger struct {
	Name     string
	LevelInt int
	Out      io.Writer
}

// NewSimpleLogger creates a new SimpleLogger with level info that writes messages to the given writer.
func NewSimpleLogger(name string, out io.Writer) *SimpleLogger {
	if out == nil {
		panic("out can not be nil")
	}

	return &SimpleLogger{
		Name:     name,
		LevelInt: LevelInfo,
		Out:      out,
	}
}

func (l *SimpleLogger) Trace(msg string, args ...interface{}) {
	l.log("TRC", msg, args...)
}

func (l *SimpleLogger) Debug(msg string, args ...interface{}) {
	l.log("DEB", msg, args...)
}

func (l *SimpleLogger) Info(msg string, args ...interface{}) {
	l.log("INF", msg, args...)
}

func (l *SimpleLogger) Warn(msg string, args ...interface{}) error {
	l.log("WRN", msg, args...)
	return fmt.Errorf(msg, args...)
}

func (l *SimpleLogger) Error(msg string, args ...interface{}) error {
	l.log("ERR", msg, args...)
	return fmt.Errorf(msg, args...)
}

func (l *SimpleLogger) log(levelString, msg string, args ...interface{}) {

	fmt.Fprintln(l.Out, time.Now().Format("15:04:05.000000"), levelString, msg, args)
}

func (l *SimpleLogger) Level() int {
	return l.LevelInt
}

func (l *SimpleLogger) SetLevel(level int) {
	l.LevelInt = level
}

func (l *SimpleLogger) IsTrace() bool {
	return l.LevelInt <= LevelTrace
}

func (l *SimpleLogger) IsDebug() bool {
	return l.LevelInt <= LevelDebug
}

func (l *SimpleLogger) IsInfo() bool {
	return l.LevelInt <= LevelInfo
}

func (l *SimpleLogger) IsWarn() bool {
	return l.LevelInt <= LevelWarn
}

func (l *SimpleLogger) IsError() bool {
	return l.LevelInt <= LevelError
}

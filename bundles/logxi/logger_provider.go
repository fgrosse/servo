package logxi

import (
	"github.com/fgrosse/servo"
	"github.com/mgutz/logxi/v1"
)

type LoggerProvider struct {
	DefaultLogLevel int
}

func (p *LoggerProvider) SetLevel(level int) {
	p.DefaultLogLevel = level
}

func (p *LoggerProvider) GetLogger(name string) servo.Logger {
	l := log.New(name)
	l.SetLevel(p.DefaultLogLevel)

	return l
}

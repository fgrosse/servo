package logxi

import (
	"github.com/mgutz/logxi/v1"
	"fmt"
)

const DefaultLevel = log.LevelWarn

type LoggerConfigurator struct {
	Level int
}

func NewLoggerConfigurator(levelString string) *LoggerConfigurator {
	level, isValid := log.LevelAtoi[levelString]
	if isValid == false {
		panic(fmt.Errorf("invalid log level %q", levelString))
	}

	return &LoggerConfigurator{level}
}

type configurable interface {
	SetLevel(int)
}

func (c *LoggerConfigurator) Configure(t configurable) {
	t.SetLevel(c.Level)
}

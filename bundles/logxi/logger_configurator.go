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

func (c *LoggerConfigurator) Configure(logger log.Logger) {
	logger.SetLevel(c.Level)
}

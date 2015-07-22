package logxi

import (
	"github.com/mgutz/logxi/v1"
	"github.com/fgrosse/servo"
)

type Bundle struct{}

func (b *Bundle) Boot(kernel *servo.Kernel) {
	kernel.Log = log.New("kernel")
}

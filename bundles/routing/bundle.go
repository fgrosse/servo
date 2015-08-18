package routing

import (
	"github.com/fgrosse/servo"
)

type Bundle struct{}

func (b *Bundle) Boot(kernel *servo.Kernel) {
	RegisterTypes(kernel.TypeRegistry)
}

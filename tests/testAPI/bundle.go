package testAPI

import "github.com/fgrosse/servo"

type TestBundle struct {}

func (b *TestBundle) Boot(kernel *servo.Kernel) {
	kernel.RegisterType("test_bundle.my_type", NewService)
}

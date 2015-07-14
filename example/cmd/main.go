package main

import (
	"github.com/fgrosse/servo"
	"github.com/fgrosse/servo/example"
)

func main() {
	kernel := servo.NewKernel("config.yml")
	example.RegisterTypes(kernel.TypeRegistry)
	kernel.Run()
}

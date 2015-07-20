package main

import (
	"github.com/fgrosse/servo"
	"github.com/fgrosse/servo/example"
	"github.com/fgrosse/servo/configuration"
)

func main() {
	config := configuration.NewYAMLFileLoader("config.yml")
	kernel := servo.NewKernel(config)
	example.RegisterTypes(kernel.TypeRegistry)
	kernel.Run()
}

package main

import (
	"github.com/fgrosse/servo"
	"github.com/fgrosse/servo/example"
)

func main() {
	config := servo.NewYAMLFileLoader("config.yml")
	kernel := servo.NewKernel(config)
	example.RegisterTypes(kernel.TypeRegistry)
	kernel.Run()
}

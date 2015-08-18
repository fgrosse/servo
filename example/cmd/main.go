package main

import (
	"log"

	"github.com/fgrosse/servo"
	"github.com/fgrosse/servo/bundles/logxi"
	"github.com/fgrosse/servo/configuration"
	"github.com/fgrosse/servo/example"
)

func main() {
	loader := configuration.NewYAMLFileLoader("config/config.yml")
	kernel := servo.NewDebugKernel(loader)
	kernel.Register(new(logxi.Bundle))

	example.RegisterTypes(kernel.TypeRegistry)
	err := kernel.Run()
	if err != nil {
		log.Fatal(err)
	}
}

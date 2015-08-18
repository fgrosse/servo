package main

import (
	"log"

	"github.com/fgrosse/servo"
	"github.com/fgrosse/servo/bundles/logxi"
	"github.com/fgrosse/servo/configuration"
	"github.com/fgrosse/servo/example"
	"github.com/fgrosse/servo/bundles/routing"
)

func main() {
	loader := configuration.NewYAMLFileLoader("config/config.yml")
	kernel := servo.NewDebugKernel(loader)
	kernel.Register(new(logxi.Bundle))
	kernel.Register(new(routing.Bundle))

	example.RegisterTypes(kernel.TypeRegistry)
	err := kernel.Run()
	if err != nil {
		log.Fatal(err)
	}
}

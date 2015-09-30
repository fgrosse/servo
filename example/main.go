package main

import (
	"log"

	"github.com/fgrosse/servo"
	"github.com/fgrosse/servo/bundles/logxi"
	"github.com/fgrosse/servo/bundles/routing"
	"github.com/fgrosse/servo/configuration"
)

func main() {
	loader := configuration.NewYAMLFileLoader("config/config.yml")
	kernel := servo.NewDebugKernel(loader)
	kernel.Register(new(logxi.Bundle))
	kernel.Register(new(routing.Bundle))

	RegisterTypes(kernel.TypeRegistry)
	err := kernel.Run()
	if err != nil {
		log.Fatal(err)
	}
}

/*
error while generating type "kernel.http.server":
	the referenced type "@kernel.http_handler"
		type *middleware.Logging
	can not be passed as argument 2 to the function signature
		servo.NewHTTPServer(string, http.HandlerFunc, servo.Logger)
 */

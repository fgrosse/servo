package servo

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/fgrosse/goldi"
	"log"
)

type Kernel struct {
	*goldi.Container
}

func NewKernel(configFilePath string) *Kernel {
	registry := goldi.NewTypeRegistry()
	registerInternalTypes(registry)

	config := loadConfiguration(configFilePath)
	return &Kernel {
		Container: goldi.NewContainer(registry, config),
	}
}

func loadConfiguration(configFilePath string) map[string]interface{} {
	// TODO load from file
	// TODO flatten configuration keys
	return map[string]interface{}{
		"servo.listen": "0.0.0.0:3000",
	}
}

func (k *Kernel) Run() {
	defer k.panicHandler()

	k.validateContainer()
	server := k.Get("kernel.server").(Server)
	err := server.Run()
	if err != nil {
		panic(err)
	}
}

func (k *Kernel) panicHandler() {
	if r := recover(); r != nil {
		// TODO implement useful panic handling
		log.Printf("PANIC: %v\n", r)
		debug.PrintStack()
		os.Exit(1)
	}
}

func (k *Kernel) validateContainer() {
	validator := goldi.NewContainerValidator()

	// TODO add explicit type checks for all internal types that might have gotten overwritten

	err := validator.Validate(k.Container)
	if err != nil {
		panic(fmt.Errorf("container validation error: %s", err))
	}
}

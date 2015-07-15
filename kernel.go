package servo

import (
	"fmt"

	"github.com/fgrosse/goldi"
)

// The Kernel is basically a goldi.TypeRegistry on which all necessary types are registered.
// Once type registration is done it is started with its Run function.
type Kernel struct {
	goldi.TypeRegistry
	Config    ConfigurationLoader
	Validator *goldi.ContainerValidator
}

// NewKernel creates a new kernel.
// It will initialize the goldi.TypeRegistry and use the given configuration loader.
// The actual loading of the configuration is deferred until kernel.Run is called.
func NewKernel(config ConfigurationLoader) *Kernel {
	kernel := &Kernel{
		TypeRegistry: goldi.NewTypeRegistry(),
		Config:       config,
		Validator:    goldi.NewContainerValidator(),
	}

	registerInternalTypes(kernel.TypeRegistry)
	return kernel
}

// Run creates a goldi.Container based on the TypeRegistry of the kernel and used the configuration loader.
// It does then instantiate the "kernel.server" type and calls Run on the resulting Server implementation.
// This method blocks until the server returns from Run.
func (k *Kernel) Run() error {
	container, err := k.createContainer()
	if err != nil {
		return err
	}

	server := container.Get("kernel.server").(Server)
	return server.Run()
}

func (k *Kernel) createContainer() (*goldi.Container, error) {
	// TODO defer panic handler for validateContainer (maybe change in goldi)
	config, err := k.Config.Load()

	if err != nil {
		return nil, err
	}

	container := goldi.NewContainer(k.TypeRegistry, config)
	k.validateContainer(container)

	return container, nil
}

func (k *Kernel) validateContainer(container *goldi.Container) {
	// TODO add explicit type checks for all internal types that might have gotten overwritten

	err := k.Validator.Validate(container)
	if err != nil {
		panic(fmt.Errorf("container validation error: %s", err))
	}
}

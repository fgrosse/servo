package servo

import (
	"github.com/fgrosse/goldi"
	"github.com/fgrosse/servo/configuration"
)

var KernelVersion = "unknown"

// The Kernel is basically a goldi.TypeRegistry on which all necessary types are registered.
// Once type registration is done it is started with its Run function.
type Kernel struct {
	goldi.TypeRegistry
	Log       Logger
	Config    ConfigurationLoader
}

// NewKernel creates a new kernel.
// It will initialize the goldi.TypeRegistry and use the given configuration loader.
// The actual loading of the configuration is deferred until kernel.Run is called.
func NewKernel(config ConfigurationLoader) *Kernel {
	kernel := &Kernel{
		TypeRegistry: goldi.NewTypeRegistry(),
		Log:          NewNullLogger(),
		Config:       config,
	}

	registerInternalTypes(kernel.TypeRegistry)
	return kernel
}

func (k *Kernel) Register(bundle Bundle) {
	bundle.Boot(k)
}

// Run creates a goldi.Container based on the TypeRegistry of the kernel and used the configuration loader.
// It does then instantiate the "kernel.server" type and calls Run on the resulting Server implementation.
// This method blocks until the server returns from Run.
func (k *Kernel) Run() error {
	k.Log.Info("Starting servo kernel", "version", KernelVersion)
	container, err := k.createContainer()
	if err != nil {
		return err
	}

	server := container.Get("kernel.server").(Server)
	return server.Run()
}

func (k *Kernel) createContainer() (*goldi.Container, error) {
	k.Log.Debug("Loading the configuration..")
	config, err := k.Config.Load()
	if err != nil {
		return nil, err
	}

	k.Log.Trace("Flattening the configuration..")
	flattenedConfig := new(configuration.Flattener).Flatten(config)
	k.Log.Debug("Configuration has been loaded", "config", flattenedConfig)

	k.Log.Debug("Creating goldi container")
	container := goldi.NewContainer(k.TypeRegistry, flattenedConfig)
	container.InjectInstance("container", container)
	container.InjectInstance("logger", k.Log)

	err = k.validateContainer(container)
	if err != nil {
		k.Log.Error("Container is invalid", "error", err)
	} else {
		k.Log.Debug("Container passed validation")
	}

	return container, err
}

const TypeContainerValidator = "container.validator"

func (k *Kernel) validateContainer(container *goldi.Container) error {
	k.Log.Trace("Retrieving validator from container", "service_name", TypeContainerValidator)
	validator := container.Get(TypeContainerValidator).(*goldi.ContainerValidator)

	k.Log.Debug("Validating container")

	return validator.Validate(container)
}

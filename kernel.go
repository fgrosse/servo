package servo

import (
	"fmt"
	"os"

	"github.com/fgrosse/goldi"
	"github.com/fgrosse/goldi/validation"
	"github.com/fgrosse/servo/configuration"
	"github.com/mgutz/logxi/v1"
)

var KernelVersion = "unknown"

// The Kernel is basically a goldi.TypeRegistry on which all necessary types are registered.
// Once type registration is done it is started with its Run function.
type Kernel struct {
	goldi.TypeRegistry

	// Config is a ConfigurationLoader that loads the configuration upon Kernel.Run
	Config ConfigurationLoader

	// Env identifies the environment the kernel runs in.
	// If Env is empty a production environment is assumed. The value is usually set via
	// the environment variable `SERVO_ENV`. Setting this variable to `dev` will configure
	// the kernel for local development (i.e. print debug messages).
	Env string

	log Logger
}

// NewKernel creates a new kernel.
// It will initialize the goldi.TypeRegistry and use the given configuration loader.
// The actual loading of the configuration is deferred until kernel.Run is called.
//
// The kernel environment identifier is read from the os environment variable
// `SERVO_ENV`. If SERVO_ENV is set to `dev` the internal log level will be set to debug.
func NewKernel(config ConfigurationLoader) *Kernel {
	kernel := &Kernel{
		TypeRegistry: goldi.NewTypeRegistry(),
		Env:          os.Getenv("SERVO_ENV"),
		Config:       config,
		log:          new(NullLogger),
	}

	if kernel.Env == "dev" {
		kernel.log = log.New("kernel")
		kernel.log.SetLevel(log.LevelDebug)
	}

	registerInternalTypes(kernel.TypeRegistry)
	return kernel
}

// Register is used to boot servo.Bundle with this kernel instance.
func (k *Kernel) Register(bundle Bundle) {
	k.log.Info("Loading bundle", "bundle", fmt.Sprintf("%T", bundle))
	bundle.Boot(k)
}

// Run load the configuration and creates the DI container.
// Afterwards the `kernel.http.server` type is created by the container and started.
func (k *Kernel) Run() error {
	k.log.Info("Starting servo kernel", "version", KernelVersion)
	container, err := k.createContainer()
	if err != nil {
		return err
	}

	server, err := container.Get("kernel.http.server")
	if err != nil {
		k.log.Fatal(err.Error())
	}

	return server.(Server).Run()
}

func (k *Kernel) createContainer() (*goldi.Container, error) {
	k.log.Debug("Loading the configuration..")
	config, err := k.Config.Load()
	if err != nil {
		return nil, err
	}

	k.log.Trace("Flattening the configuration..")
	flattenedConfig := new(configuration.Flattener).Flatten(config)
	k.log.Debug("Configuration has been loaded", "config", flattenedConfig)

	k.log.Debug("Creating goldi container")
	container := goldi.NewContainer(k.TypeRegistry, flattenedConfig)
	container.InjectInstance("container", container)

	err = k.validateContainer(container)
	if err != nil {
		k.log.Error("Container is invalid", "error", err)
	} else {
		k.log.Debug("Container passed validation")
	}

	return container, err
}

// TypeContainerValidator is the type ID that is used to get the container validator from the DI container.
const TypeContainerValidator = "container.validator"

func (k *Kernel) validateContainer(container *goldi.Container) error {
	k.log.Trace("Retrieving validator from container", "service_name", TypeContainerValidator)
	validator, err := container.Get(TypeContainerValidator)
	if err != nil {
		k.log.Fatal(err.Error())
	}

	k.log.Debug("Validating container")

	return validator.(*validation.ContainerValidator).Validate(container)
}

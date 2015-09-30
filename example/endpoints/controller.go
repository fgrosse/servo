package endpoints

import (
	"github.com/fgrosse/servo/example/lib"
	"fmt"
	"net/http"
	"github.com/fgrosse/goldi"
	"github.com/fgrosse/servo"
)

type FancyController struct {
	Client lib.ServiceClient
}

func (c *FancyController) FancyAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service said: %q", c.Client.DoFancyStuff())
}

func (c *FancyController) OuterHandlerAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OuterHandlerAction was called")
}

func (c *FancyController) SecondHandlerAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "SecondHandlerAction was called")
}

func (c *FancyController) ErrorAction(w http.ResponseWriter, r *http.Request) {
	panic(fmt.Errorf("OH MY GOD!"))
}

type ContainerAwareController struct {
	Container *goldi.Container
}

func (c *ContainerAwareController) SomeAction(w http.ResponseWriter, r *http.Request) {
	logger := c.Container.MustGet("logger").(servo.Logger)
	logger.Info("Logger was called from within ContainerAwareController")

	fmt.Fprintf(w, "Wrote message to servo logger")
}

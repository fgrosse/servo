package endpoints

import (
	"github.com/fgrosse/servo/example/lib"
	"fmt"
	"net/http"
)

type FancyController struct {
	Client lib.ServiceClient
}

func (c *FancyController) FancyAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service said: %q", c.Client.DoFancyStuff())
}

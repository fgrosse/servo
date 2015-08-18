//go:generate goldigen --in "servo_types.yml" --out "servo_types.go" --package github.com/fgrosse/servo --function registerInternalTypes --overwrite --nointeraction
package servo

import (
	"github.com/fgrosse/goldi"
)

// registerInternalTypes registers all types that have been defined in the file "servo_types.yml"
//
// DO NOT EDIT THIS FILE: it has been generated by goldigen v0.9.0.
// It is however good practice to put this file under version control.
// See https://github.com/fgrosse/goldi for what is going on here.
func registerInternalTypes(types goldi.TypeRegistry) {
	types.Register("container.validator", goldi.NewType(goldi.NewContainerValidator))
	types.Register("kernel.server", goldi.NewType(NewHTTPServer, "%servo.listen%", "@kernel.http_handler", "@logger"))
	types.Register("logger", goldi.NewType(NewNullLogger))
}

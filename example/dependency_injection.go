package example

import (
	"github.com/fgrosse/goldi"
)

// RegisterTypes registers all types that have been defined in the file "types.yml"
//
// DO NOT EDIT THIS FILE: it has been generated by goldigen v0.1.0.
// It is however good practice to put this file under version control.
// See https://github.com/FGrosse/goldi for what is going on here.
func RegisterTypes(types goldi.TypeRegistry) {
	types.RegisterType("kernel.http_handler", NewMySimpleHandler)
}

package example

import "net/http"

type MySimpleHandler struct {}

func NewMySimpleHandler() *MySimpleHandler {
	return &MySimpleHandler{}
}

func (c *MySimpleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello servo world :)"))
}

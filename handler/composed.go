package handler

import "net/http"

type composed struct {
	handlers []http.HandlerFunc
}

// Composed is used to compose an http.Handler from several http.HandlerFunc instances.
// The handler functions are called in the order in which they are passed to this function.
func Composed(handlers ...http.HandlerFunc) *composed {
	return &composed{handlers}
}

// ServeHTTP implements the http.Handler interface by calling all registered handler functions in order.
func (c *composed) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, h := range c.handlers {
		h(w, r)
	}
}

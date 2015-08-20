package handlers

import "net/http"

type composedHandler struct {
	handlers []http.HandlerFunc
}

func ComposedHandler(handlers ...http.HandlerFunc) *composedHandler {
	return &composedHandler{handlers}
}

func (c *composedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, h := range c.handlers {
		h(w, r)
	}
}

package handlers

import "net/http"

// The Adapter follows an idea for designing middleware in go
// See https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
type Adapter func(http.HandlerFunc) http.HandlerFunc

type adapterHandler struct {
	handlerFunc http.HandlerFunc
}

func AdapterHandler(adapters ...Adapter) http.Handler {
	h := &adapterHandler{}

	var handlerFunc http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {}
	for _, adapter := range adapters {
		handlerFunc = adapter(handlerFunc)
	}

	h.handlerFunc = handlerFunc
	return h
}

func (h *adapterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handlerFunc(w, r)
}

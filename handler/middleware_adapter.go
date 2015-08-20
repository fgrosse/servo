package handler

import "net/http"

// A Middleware is a function that creates http handlers from other http handlers.
// Usually the corresponding middleware creates a new handler that will call the given handler at some point.
// See https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81
type Middleware func(http.Handler) http.Handler

func MiddleWareAdapter(h http.Handler, adapters ...Middleware) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}

	return h
}

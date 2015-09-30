package middleware

import (
	"net/http"

	"github.com/fgrosse/servo/handler"
)

type ErrorHandlingFunc func(recovered interface{}, w http.ResponseWriter, r *http.Request)

type ErrorHandling struct {
	http.Handler
	ErrorHandler ErrorHandlingFunc
}

// ErrorHandlingAdapter creates a new ErrorHandling middleware
func ErrorHandlingAdapter(errorHandler ErrorHandlingFunc) handler.Middleware {
	return func(h http.Handler) http.Handler {
		return &ErrorHandling{h, errorHandler}
	}
}

func (h *ErrorHandling) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer h.panicHandler(w, r)
	h.Handler.ServeHTTP(w, r)
}

func (h *ErrorHandling) panicHandler(w http.ResponseWriter, rq *http.Request) {
	if r := recover(); r != nil {
		h.ErrorHandler(r, w, rq)
	}
}

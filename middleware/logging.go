package middleware

import (
	"net/http"

	"github.com/fgrosse/servo"
	"github.com/fgrosse/servo/handler"
)

// LoggingMiddleWare implements the http.Handler interface by decorating another
// http handler with logging functionality.
// The middleware will log requests and responses both without body.
type Logging struct {
	http.Handler
	Logger servo.Logger
}

// LoggingAdapter creates a new LoggingMiddleWare
func LoggingAdapter(logger servo.Logger) handler.Middleware {
	return func(h http.Handler) http.Handler {
		return &Logging{Handler: h, Logger: logger}
	}
}

func (m *Logging) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var spy *responseSpy
	if m.Logger.IsInfo() {
		m.Logger.Info("Serving request",
			"method", r.Method,
			"uri", r.RequestURI,
			"headers", r.Header,
			"content-length", r.ContentLength,
			"remote", r.RemoteAddr,
			"host", r.Host,
		)
		spy = &responseSpy{response: w}
		w = spy
	}

	defer func() {
		if m.Logger.IsInfo() {
			m.Logger.Info("Finished request",
				"method", r.Method,
				"url", r.RequestURI,
				"status", spy.servedStatus(),
				"remote", r.RemoteAddr,
			)
		}
	}()

	m.Handler.ServeHTTP(w, r)
}

type responseSpy struct {
	response http.ResponseWriter
	status   int
}

func (r *responseSpy) Header() http.Header {
	return r.response.Header()
}

func (r *responseSpy) Write(data []byte) (int, error) {
	return r.response.Write(data)
}

func (r *responseSpy) WriteHeader(status int) {
	r.status = status
	r.response.WriteHeader(status)
}

func (r *responseSpy) servedStatus() int {
	if r.status == 0 {
		return http.StatusOK
	}

	return r.status
}

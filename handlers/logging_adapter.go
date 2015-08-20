package handlers

import (
	"net/http"

	"github.com/fgrosse/servo"
)

func LoggingAdapter(logger servo.Logger) Adapter {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var spy *responseSpy
			if logger.IsInfo() {
				logger.Info("Serving request",
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
				if logger.IsInfo() {
					logger.Info("Finished request",
						"method", r.Method,
						"url", r.RequestURI,
						"status", spy.servedStatus(),
						"remote", r.RemoteAddr,
					)
				}
			}()

			h(w, r)
		}
	}
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

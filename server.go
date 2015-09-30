package servo

import "net/http"

// The Server is responsible for running the server process.
// It can be injected via the goldi type "kernel.server"
type Server interface {
	// Run starts the server and blocks until it has finished
	Run() error
}

// DefaultServer is the standard implementation of the Server interface.
type HTTPServer struct {
	ListenAddress string
	HandlerFunc   http.HandlerFunc
	Log           Logger
}

// NewHTTPServer creates a new HTTPServer
func NewHTTPServer(listenAddress string, handler http.HandlerFunc, log Logger) *HTTPServer {
	return &HTTPServer{listenAddress, handler, log}
}

// Run will make this server listen on the given ListenAddress and use the handler to
// handle all incoming HTTP requests. The method blocks.
func (s *HTTPServer) Run() error {
	s.Log.Info("Server started", "address", s.ListenAddress)
	http.Handle("/", s.Handler)
	return http.ListenAndServe(s.ListenAddress, nil)
}

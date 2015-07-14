package servo

import "net/http"

type Server interface {
	Run() error
}

type DefaultServer struct {
	ListenAddress string
	Handler       http.Handler
}

func NewDefaultServer(listenAddress string, handler http.Handler) *DefaultServer {
	return &DefaultServer{listenAddress, handler}
}

func (s *DefaultServer) Run() error {
	return http.ListenAndServe(s.ListenAddress, s.Handler)
}

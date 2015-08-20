package routing

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/fgrosse/goldi"
	"fmt"
)

type Router struct {
	*mux.Router
}

func NewRouter(loader *Loader, container *goldi.Container) *Router {
	routes, err := loader.Load("config/routes.yml")
	if err != nil {
		panic(err)
	}

	r := &Router{
		Router: mux.NewRouter(),
	}

	for _, route := range routes {
		handler, err := handlerFunc(route.EndpointTypeID, container)
		if err != nil {
			panic(err) // TODO implement proper error handling (teach goldi that it is ok to return an error as last return argument from a factory
		}
		r.HandleFunc(route.Path, handler)
	}

	return r
}

func handlerFunc(endpointTypeID string, container *goldi.Container) (http.HandlerFunc, error) {
	h, err := container.Get(endpointTypeID)
	if err != nil {
		return nil, err
	}

	if h, ok := h.(func(http.ResponseWriter, *http.Request)); ok {
		return h, nil
	}

	if h, ok := h.(http.Handler); ok {
		return func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}, nil
	}

	return nil, fmt.Errorf("Type %q is neither a http.Handler not a http.HandlerFunc but a %T", endpointTypeID, h)
}

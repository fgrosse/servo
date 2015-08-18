package routing

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/fgrosse/goldi"
)

func NewRouter(loader *Loader, container *goldi.Container) http.Handler {
	routes, err := loader.Load("config/routes.yml")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	for _, route := range routes {
		endpointTypeID := route.EndpointTypeID[1:]
		handler := container.Get(endpointTypeID).(func(http.ResponseWriter, *http.Request)) // TODO check if endpointTypeID is a valid type ID
		r.HandleFunc(route.Path, handler)
	}

	return r
}

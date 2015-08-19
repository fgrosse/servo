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
		handler := container.Get(route.EndpointTypeID).(func(http.ResponseWriter, *http.Request))
		r.HandleFunc(route.Path, handler)
	}

	return r
}

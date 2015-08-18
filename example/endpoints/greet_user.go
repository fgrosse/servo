package endpoints

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func GreetUserEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello %s!", name)
}

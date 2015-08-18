package endpoints

import "net/http"

func HelloWorldEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello servo world :)"))
}

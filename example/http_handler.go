package example

import "net/http"

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello servo world :)"))
}

package lib

import (
	"net/http"
	"github.com/fgrosse/servo"
	"fmt"
)

type MyErrorHandler struct {
	Logger servo.Logger
}

func (h *MyErrorHandler) HandleEndpointError(recovered interface{}, w http.ResponseWriter, r *http.Request) {
	if h.Logger != nil {
		h.Logger.Error("Oh no!", "error", recovered, "request", r, "response")
	}
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Sorry but something went horribly wrong: %s", recovered)
}

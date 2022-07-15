package http2

import (
	"net/http"
)

var NotFound func(http.ResponseWriter, *http.Request)

func Error(w http.ResponseWriter, r *http.Request) {
	if NotFound == nil {
		http.NotFound(w, r)
		return
	}
	NotFound(w, r)
}

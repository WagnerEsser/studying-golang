package handlers

import (
	"net/http"
	restError "studying-go/types"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		restError.NewMethodNotAllowedError().Throw(w)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, " + name + "!"))
}

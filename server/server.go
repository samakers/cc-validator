package server

import (
	"net/http"
)

func Start() {
	http.HandleFunc("/validate", ValidateHandler)
	http.ListenAndServe(":8090", nil)
}

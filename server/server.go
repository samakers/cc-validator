package server

import (
	"net/http"
)

func Start() {
	http.HandleFunc("/validate", splitInteger)
	http.ListenAndServe(":8090", nil)
}

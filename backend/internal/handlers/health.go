package handlers

import (
	"net/http"
)

// HealthCheckHandler responde se o serviço está online
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

package main

import (
	"fmt"
	"net/http"
	"sistemadefila/backend/internal/handlers"
)

func main() {
	fmt.Println("Servidor de fila iniciado na porta 8080...")
	http.HandleFunc("/health", handlers.HealthCheckHandler)
	http.HandleFunc("/enqueue", handlers.EnqueueHandler)
	http.HandleFunc("/dequeue", handlers.DequeueHandler)
	http.ListenAndServe(":8080", nil)
}

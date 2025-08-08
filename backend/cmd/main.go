package main

import (
	"fmt"
	"net/http"
	_ "sistemadefila/backend/docs"
	"sistemadefila/backend/internal/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	fmt.Println("Servidor de fila iniciado na porta 8080...")
	http.HandleFunc("/health", handlers.HealthCheckHandler)
	http.HandleFunc("/enqueue", handlers.EnqueueHandler)
	http.HandleFunc("/dequeue", handlers.DequeueHandler)
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.ListenAndServe(":8080", nil)
}

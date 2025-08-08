package handlers

import (
	"encoding/json"
	"net/http"
	"sistemadefila/backend/internal/models"
	"sync"
)

var (
	queue     = models.NewQueue()
	queueLock sync.Mutex
)

func EnqueueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Item string `json:"item"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Item == "" {
		http.Error(w, "JSON inválido ou item vazio", http.StatusBadRequest)
		return
	}
	queueLock.Lock()
	queue.Enqueue(req.Item)
	queueLock.Unlock()
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("adicionado"))
}

func DequeueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	queueLock.Lock()
	item, ok := queue.Dequeue()
	queueLock.Unlock()
	if !ok {
		http.Error(w, "Fila vazia", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"item": item})
}

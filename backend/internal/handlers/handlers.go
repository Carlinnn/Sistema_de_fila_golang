package handlers

import (
	"encoding/json"
	"net/http"
	"sistemadefila/backend/internal/models"
	"sync"
)

var (
	queues     = make(map[string]*models.Queue)
	queuesLock sync.Mutex
)

func EnqueueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Queue string `json:"queue"`
		Item  string `json:"item"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Item == "" || req.Queue == "" {
		http.Error(w, "JSON inválido ou campos obrigatórios ausentes", http.StatusBadRequest)
		return
	}
	queuesLock.Lock()
	q, ok := queues[req.Queue]
	if !ok {
		q = models.NewQueue()
		queues[req.Queue] = q
	}
	q.Enqueue(req.Item)
	queuesLock.Unlock()
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("adicionado"))
}

func DequeueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Queue string `json:"queue"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Queue == "" {
		http.Error(w, "JSON inválido ou campo 'queue' ausente", http.StatusBadRequest)
		return
	}
	queuesLock.Lock()
	q, ok := queues[req.Queue]
	if !ok {
		queuesLock.Unlock()
		http.Error(w, "Fila não encontrada", http.StatusNotFound)
		return
	}
	item, ok := q.Dequeue()
	queuesLock.Unlock()
	if !ok {
		http.Error(w, "Fila vazia", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"item": item})
}

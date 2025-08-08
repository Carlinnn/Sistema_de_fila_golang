package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sistemadefila/backend/internal/models"
	"testing"
)

func TestEnqueueDequeueHandlers(t *testing.T) {
	queuesLock.Lock()
	queues = make(map[string]*models.Queue)
	queuesLock.Unlock()

	body := bytes.NewBufferString(`{"queue":"fila1","item":"teste1"}`)
	req := httptest.NewRequest("POST", "/enqueue", body)
	rr := httptest.NewRecorder()
	EnqueueHandler(rr, req)
	if rr.Code != http.StatusCreated {
		t.Errorf("esperado 201, obteve %d", rr.Code)
	}

	body = bytes.NewBufferString(`{"queue":"fila1"}`)
	req = httptest.NewRequest("POST", "/dequeue", body)
	rr = httptest.NewRecorder()
	DequeueHandler(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("esperado 200, obteve %d", rr.Code)
	}
	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil || resp["item"] != "teste1" {
		t.Errorf("esperado item 'teste1', obteve '%v' (err=%v)", resp, err)
	}

	body = bytes.NewBufferString(`{"queue":"fila1"}`)
	req = httptest.NewRequest("POST", "/dequeue", body)
	rr = httptest.NewRecorder()
	DequeueHandler(rr, req)
	if rr.Code != http.StatusNotFound {
		t.Errorf("esperado 404 para fila vazia, obteve %d", rr.Code)
	}

	// Teste fila inexistente
	body = bytes.NewBufferString(`{"queue":"fila2"}`)
	req = httptest.NewRequest("POST", "/dequeue", body)
	rr = httptest.NewRecorder()
	DequeueHandler(rr, req)
	if rr.Code != http.StatusNotFound {
		t.Errorf("esperado 404 para fila inexistente, obteve %d", rr.Code)
	}
}

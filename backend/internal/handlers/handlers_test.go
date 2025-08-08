package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEnqueueDequeueHandlers(t *testing.T) {
	queueLock.Lock()
	queue.Items = []string{}
	queueLock.Unlock()

	body := bytes.NewBufferString(`{"item":"teste1"}`)
	req := httptest.NewRequest("POST", "/enqueue", body)
	rr := httptest.NewRecorder()
	EnqueueHandler(rr, req)
	if rr.Code != http.StatusCreated {
		t.Errorf("esperado 201, obteve %d", rr.Code)
	}

	req = httptest.NewRequest("POST", "/dequeue", nil)
	rr = httptest.NewRecorder()
	DequeueHandler(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("esperado 200, obteve %d", rr.Code)
	}
	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil || resp["item"] != "teste1" {
		t.Errorf("esperado item 'teste1', obteve '%v' (err=%v)", resp, err)
	}

	rr = httptest.NewRecorder()
	DequeueHandler(rr, req)
	if rr.Code != http.StatusNotFound {
		t.Errorf("esperado 404 para fila vazia, obteve %d", rr.Code)
	}
}

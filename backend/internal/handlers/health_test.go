package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()

	HealthCheckHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", rr.Code)
	}
	if rr.Body.String() != "ok" {
		t.Errorf("esperado corpo 'ok', obteve '%s'", rr.Body.String())
	}
}

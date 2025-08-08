package models

import "testing"

func TestQueue_BasicFlow(t *testing.T) {
	q := NewQueue()
	if q.Len() != 0 {
		t.Errorf("esperado fila vazia, obteve %d", q.Len())
	}

	q.Enqueue("primeiro")
	q.Enqueue("segundo")
	if q.Len() != 2 {
		t.Errorf("esperado 2 itens, obteve %d", q.Len())
	}

	item, ok := q.Dequeue()
	if !ok || item != "primeiro" {
		t.Errorf("esperado 'primeiro', obteve '%s' (ok=%v)", item, ok)
	}

	item, ok = q.Dequeue()
	if !ok || item != "segundo" {
		t.Errorf("esperado 'segundo', obteve '%s' (ok=%v)", item, ok)
	}

	if q.Len() != 0 {
		t.Errorf("esperado fila vazia após remoção, obteve %d", q.Len())
	}

	_, ok = q.Dequeue()
	if ok {
		t.Errorf("esperado dequeue falso em fila vazia")
	}
}
